from locust import HttpUser, task

class QuickstartUser(HttpUser):
    @task
    def auth(self):
        self.client.post(url="/v1/auth", headers={"Content-Type": "application/json"}, json={"email": "aaa@example.com", "password": "passwd"})