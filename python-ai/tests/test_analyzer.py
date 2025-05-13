# tests/test_analyzer.py
import unittest
from unittest.mock import AsyncMock

from app.analyzer import analyze_conversation
from app.main import AnalyzeResponse

class TestAnalyzeConversation(unittest.IsolatedAsyncioTestCase):
    async def test_analyze_conversation(self):
        mock_chain = AsyncMock()
        mock_chain.ainvoke.return_value = '{"issues": ["Ambiguous response"], "suggestions": ["Clarify intent"]}'
        
        result: AnalyzeResponse = await analyze_conversation("Speaker A: Hi\nSpeaker B: What?", mock_chain)

        self.assertIsInstance(result, AnalyzeResponse)
        self.assertEqual(result.issues, ["Ambiguous response"])
        self.assertEqual(result.suggestions, ["Clarify intent"])
        
    async def test_invalid_json_response_raises_exception(self):
        mock_chain = AsyncMock()
        mock_chain.ainvoke.return_value = 'this is not json'
        
        from fastapi import HTTPException
        with self.assertRaises(HTTPException) as context:
            await analyze_conversation("Speaker A: Hello\nSpeaker B: ???", mock_chain)
        self.assertEqual(context.exception.status_code, 500)
        self.assertIn("Invalid JSON format", context.exception.detail)

if __name__ == "__main__":
    unittest.main()
