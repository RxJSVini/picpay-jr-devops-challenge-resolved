FROM python:3.9

WORKDIR /usr/src/app

COPY requirements.txt  .
COPY main.py .

RUN pip install -r requirements.txt

EXPOSE 8081

CMD ["python", "main.py"]
