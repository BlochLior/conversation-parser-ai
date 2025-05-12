import unittest
from unittest.mock import patch, AsyncMock

from app.analyzer import analyze_conversation
from app.main import AnalyzeResponse
from fastapi import HTTPException

class TestAnalyzeConversation(unittest.IsolatedAsyncioTestCase):
    async def test_analyze_conversation(self):
        mock_result = '{"issues": ["Ambiguous response"], "suggestions": ["Clarify intent"]}'
        
        with patch("app.analyzer.chain", new=AsyncMock()) as mock_chain:
            mock_chain.ainvoke.return_value = mock_result
            
            result: AnalyzeResponse = await analyze_conversation("Speaker A: Hi \nSpeaker B: What?")
            
            self.assertIsInstance(result, AnalyzeResponse)
            self.assertEqual(result.issues, ["Ambiguous response"])
            self.assertEqual(result.suggestions, ["Clarify intent"])
            
    async def test_invalid_json_response_raises_exception(self):
        invalid_json = 'this is not a json'
        
        with patch("app.analyzer.chain", new=AsyncMock()) as mock_chain:
            mock_chain.ainvoke.return_value = invalid_json
            
            with self.assertRaises(HTTPException) as context:
                await analyze_conversation("Speaker A: Hello\nSpeaker B: ???")
            
            self.assertEqual(context.exception.status_code, 500)
            self.assertIn("Invalid JSON format", context.exception.detail)
            
if __name__ == "__main__":
    unittest.main()