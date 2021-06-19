
FROM python
WORKDIR /dockerapp

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY . .
EXPOSE 5000
CMD ["python3","./app.py"]
