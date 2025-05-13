from fastapi import FastAPI
from app.analyzer import analyze_conversation, ChatOpenAI, prompt, StrOutputParser, model_name
from app.schema import AnalyzeRequest, AnalyzeResponse

app = FastAPI()

# Construct chain once per app lifecycle
llm = ChatOpenAI(temperature=0.7, model=model_name)
chain = prompt | llm | StrOutputParser()

@app.post("/analyze", response_model=AnalyzeResponse)
async def analyze(req: AnalyzeRequest):
    return await analyze_conversation(req.conversation, chain)


