import json
from flask import Flask, render_template,request,session

from requests.auth import HTTPBasicAuth

def test_index(app, client):
    del app
    res = client.get('/')
    assert res.status_code == 200

def test_signup(app, client):
    del app
    res = client.get('/signup')
    assert res.status_code == 200

    