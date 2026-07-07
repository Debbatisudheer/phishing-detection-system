import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../services/api";

function Login() {

  const navigate = useNavigate();

  const [username, setUsername] =
    useState("");

  const [password, setPassword] =
    useState("");

  const login = async () => {

    try {

      const response =
        await api.post(
          "/api/login",
          {
            username,
            password,
          },
        );

      localStorage.setItem(
        "token",
        response.data.token,
      );

      alert(
        "Login Success",
      );

      navigate(
        "/",
        {
          replace: true,
        },
      );

    } catch (error) {

      console.error(error);

      alert(
        "Invalid Username or Password",
      );

    }

  };

  return (

    <div>

      <h1>

        Login

      </h1>

      <input
        placeholder="Username"
        value={username}
        onChange={(e) =>
          setUsername(
            e.target.value,
          )
        }
      />

      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) =>
          setPassword(
            e.target.value,
          )
        }
      />

      <button
        onClick={login}
      >

        Login

      </button>

    </div>

  );

}

export default Login;