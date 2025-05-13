from pydantic import BaseModel
from typing import List

class AnalyzeRequest(BaseModel):
    conversation: str
    
class AnalyzeResponse(BaseModel):
    issues: List[str]
    suggestions: List[str]