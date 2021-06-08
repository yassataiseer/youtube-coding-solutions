



def test_index(app, client):
    del app
    status = client.get("/")
    assert status.status_code == 200

def test_signup(app, client):
    del app
    status = client.get("/signup")
    assert status.status_code == 200