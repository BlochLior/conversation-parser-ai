from fastapi import FastAPI
from pydantic import BaseModel
from typing import List
from app.analyzer import analyze_conversation

app = FastAPI()

class AnalyzeRequest(BaseModel):
    conversation: str
    
class AnalyzeResponse(BaseModel):
    issues: List[str]
    suggestions: List[str]
    
@app.post("/analyze", response_model=AnalyzeResponse)
async def analyze(req: AnalyzeRequest):
    return await analyze_conversation(req.conversation)


