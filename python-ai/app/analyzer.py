import os
import json
from typing import List
from app.schema import AnalyzeResponse
from langchain_core.prompts import PromptTemplate
from langchain_core.output_parsers import StrOutputParser
from langchain_openai import ChatOpenAI
from fastapi import HTTPException
from pathlib import Path
import logging

logger = logging.getLogger("uvicorn")

if os.getenv("ENV") != "production":
    from dotenv import load_dotenv
    load_dotenv(dotenv_path=Path(__file__).resolve().parents[2] / ".env")
    
model_name = os.getenv("OPENAI_MODEL", "gpt-3.5-turbo")

prompt = PromptTemplate(
    input_variables=["conversation"],
    template="""
You are a communication coach. Analyze the following conversation.
Identify any communication issues or misunderstandings, and suggest improvements.

Conversation:
{conversation}

Return a JSON with \"issues\" and \"suggestions\" as lists.
"""
)

def safe_parse_json(response_text: str) -> dict:
    try:
        return json.loads(response_text)
    except json.JSONDecodeError as e:
        logger.error(f"❌ Failed to parse JSON: {e}\nResponse text: {response_text}")
        raise HTTPException(status_code=500, detail="Invalid JSON format returned by AI")

async def analyze_conversation(text: str, chain) -> AnalyzeResponse:
    try:
        result = await chain.ainvoke({"conversation": text})
        parsed = safe_parse_json(result)
        return AnalyzeResponse(
            issues=parsed.get("issues", []),
            suggestions=parsed.get("suggestions", [])
        )
    except Exception as e:
        logger.exception((f"❌ AI analysis failed: {e}"))
        raise HTTPException(status_code=500, detail="AI analysis failed")