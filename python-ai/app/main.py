from fastapi import FastAPI
from app.analyzer import analyze_conversation, ChatOpenAI, prompt, StrOutputParser, model_name
from app.schema import AnalyzeRequest, AnalyzeResponse
import logging
import os

# Ensure log dir exists
os.makedirs("logs", exist_ok=True)

logging.basicConfig(
    format="%(asctime)s [%(levelname)s] %(message)s",
    level=logging.INFO,
    handlers=[
        logging.StreamHandler(),
        logging.FileHandler("logs/app.log")
    ]
)
logger = logging.getLogger("uvicorn")

app = FastAPI()

llm = ChatOpenAI(temperature=0.7, model=model_name)
chain = prompt | llm | StrOutputParser()

@app.get("/health")
async def health_check():
    return {"status": "ok"}

@app.post("/analyze", response_model=AnalyzeResponse)
async def analyze(req: AnalyzeRequest):
    logger.info(f"🔍 Received conversation:\n{req.conversation}")
    result = await analyze_conversation(req.conversation, chain)
    logger.info(f"✅ AI response: {result}")
    return result


