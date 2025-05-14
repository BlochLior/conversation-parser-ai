import os
import json
from typing import List
from app.schema import AnalyzeResponse
from langchain_core.prompts import PromptTemplate
from langchain_core.output_parsers import StrOutputParser
from langchain_openai import ChatOpenAI
from dotenv import load_dotenv
from fastapi import HTTPException
from pathlib import Path

if os.getenv("ENV") != "production":
    load_dotenv(dotenv_path=Path(__file__).resolve().parents[2] / ".env")
    
model_name = os.getenv("OPENAI_MODEL", "gpt-3.5-turbo")

llm = ChatOpenAI(temperature=0.7, model=model_name)

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

chain = prompt | llm | StrOutputParser()

def safe_parse_json(response_text: str) -> dict:
    try:
        return json.loads(response_text)
    except json.JSONDecodeError:
        raise HTTPException(status_code=500, detail="Invalid JSON format returned by AI")

async def analyze_conversation(text: str, chain) -> AnalyzeResponse:
    result = await chain.ainvoke({"conversation": text})
    parsed = safe_parse_json(result)
    return AnalyzeResponse(
        issues=parsed.get("issues", []),
        suggestions=parsed.get("suggestions", [])
    )
