import "./App.css";
import { useState } from "react";
import axios from "axios";

function App() {
  const [email, setEmail] = useState("aaa@example.com");
  const [password, setPassword] = useState("passwd");
  const [authResult, setAuthResult] = useState("");
  const [token, setToken] = useState("");
  const [verifyResult, setVerifyResult] = useState("");

  const auth = () => {
    axios
      .post("http://localhost/v1/auth", {
        email: email,
        password: password,
      })
      .then((res) => {
        setAuthResult(JSON.stringify(res.data));
        setToken(res.data.token);
      })
      .catch((e) => {
        setAuthResult("エラーが発生しました");
      });
  };

  const verify = () => {
    axios
      .post("http://localhost/v1/verify", {
        token: token,
      })
      .then((res) => {
        setVerifyResult(JSON.stringify(res.data));
      })
      .catch(() => {
        setVerifyResult("エラーが発生しました");
      });
  };

  return (
    <div className="App">
      <label>
        email:
        <input
          type="text"
          defaultValue={email}
          onChange={(e) => setEmail(e.target.value)}
        ></input>
      </label>
      <label>
        password:
        <input
          type="text"
          defaultValue={password}
          onChange={(e) => setPassword(e.target.value)}
        ></input>
      </label>
      <button onClick={auth}>auth</button>
      <div>auth:{authResult}</div>
      <label>
        token:
        <input
          type="text"
          defaultValue={token}
          onChange={(e) => setToken(e.target.value)}
        ></input>
      </label>
      <button onClick={verify}>verify</button>
      <div>verify:{verifyResult}</div>
    </div>
  );
}

export default App;
