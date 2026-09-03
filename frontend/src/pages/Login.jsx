import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../services/api";

function Login() {
  const navigate = useNavigate();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const login = async () => {
    try {
      const response = await api.post("/api/login", {
        username,
        password,
      });

      localStorage.setItem(
        "token",
        response.data.token
      );

      alert("Login Success");

      navigate("/", {
        replace: true,
      });
    } catch (error) {
      console.error(error);

      alert("Invalid Username or Password");
    }
  };

  return (
    <div className="login-page">
      {/* Background */}
      <div className="login-background" aria-hidden="true">
        <div className="login-glow login-glow-one" />
        <div className="login-glow login-glow-two" />
        <div className="login-grid-pattern" />
      </div>

      {/* Login Container */}
      <div className="login-container">
        <div className="login-card">
          {/* Header */}
          <div className="login-header">
            <div className="login-brand">
              <span className="login-brand-primary">
                SOC
              </span>
              <span className="login-brand-accent">
                .Platform
              </span>
            </div>

            <div className="login-status">
              <span className="login-status-dot" />
              Secure Access
            </div>

            <h1>Login</h1>

            <p>
              Sign in to access the Security Operations
              Center dashboard.
            </p>
          </div>

          {/* Form */}
          <div className="login-form">
            {/* Username */}
            <div className="login-field">
              <label htmlFor="username">
                Username
              </label>

              <input
                id="username"
                type="text"
                placeholder="Enter your username"
                value={username}
                onChange={(e) =>
                  setUsername(e.target.value)
                }
                autoComplete="username"
              />
            </div>

            {/* Password */}
            <div className="login-field">
              <div className="login-password-header">
                <label htmlFor="password">
                  Password
                </label>
              </div>

              <input
                id="password"
                type="password"
                placeholder="Enter your password"
                value={password}
                onChange={(e) =>
                  setPassword(e.target.value)
                }
                autoComplete="current-password"
              />
            </div>

            {/* Login Button */}
            <button
              type="button"
              onClick={login}
              className="login-button"
            >
              <span>Login</span>
              <span className="login-button-arrow">
                →
              </span>
            </button>
          </div>

          {/* Security Notice */}
          <div className="login-security">
            <span className="login-security-icon">
              🔒
            </span>

            <div>
              <strong>Protected Access</strong>

              <p>
                Authorized users only. Your session is
                protected by secure authentication.
              </p>
            </div>
          </div>
        </div>

        {/* Footer */}
        <div className="login-footer">
          <span>Security Operations Center</span>

          <span className="login-footer-separator">
            •
          </span>

          <span>Secure Authentication</span>
        </div>
      </div>

      {/* Styles */}
      <style>
        {`
          .login-page {
            position: relative;
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 32px 20px;
            overflow: hidden;
            box-sizing: border-box;
            background:
              radial-gradient(
                circle at top,
                #12233d 0%,
                #081525 45%,
                #040b14 100%
              );
            color: #e6edf7;
          }

          .login-background {
            position: absolute;
            inset: 0;
            pointer-events: none;
            overflow: hidden;
          }

          .login-glow {
            position: absolute;
            border-radius: 50%;
            filter: blur(100px);
          }

          .login-glow-one {
            width: 380px;
            height: 380px;
            top: -160px;
            left: -100px;
            background: rgba(0, 196, 159, 0.08);
          }

          .login-glow-two {
            width: 420px;
            height: 420px;
            right: -180px;
            bottom: -180px;
            background: rgba(59, 130, 246, 0.08);
          }

          .login-grid-pattern {
            position: absolute;
            inset: 0;
            opacity: 0.12;
            background-image:
              linear-gradient(
                rgba(255, 255, 255, 0.035) 1px,
                transparent 1px
              ),
              linear-gradient(
                90deg,
                rgba(255, 255, 255, 0.035) 1px,
                transparent 1px
              );
            background-size: 48px 48px;
          }

          .login-container {
            position: relative;
            z-index: 1;
            width: 100%;
            max-width: 440px;
          }

          .login-card {
            width: 100%;
            padding: 38px;
            border: 1px solid #20344e;
            border-radius: 18px;
            background:
              linear-gradient(
                145deg,
                rgba(13, 29, 48, 0.97),
                rgba(7, 18, 32, 0.98)
              );
            box-shadow:
              0 30px 80px rgba(0, 0, 0, 0.38),
              0 0 0 1px rgba(255, 255, 255, 0.015);
            backdrop-filter: blur(20px);
            box-sizing: border-box;
          }

          .login-header {
            text-align: center;
          }

          .login-brand {
            display: inline-flex;
            align-items: center;
            margin-bottom: 18px;
            font-size: 18px;
            font-weight: 700;
            letter-spacing: 0.02em;
          }

          .login-brand-primary {
            color: #f4f7fb;
          }

          .login-brand-accent {
            color: #00c49f;
          }

          .login-status {
            display: inline-flex;
            align-items: center;
            gap: 7px;
            padding: 6px 10px;
            margin-bottom: 18px;
            border: 1px solid rgba(0, 196, 159, 0.18);
            border-radius: 999px;
            background: rgba(0, 196, 159, 0.05);
            color: #7f9baf;
            font-size: 11px;
            font-weight: 600;
            letter-spacing: 0.04em;
          }

          .login-status-dot {
            width: 6px;
            height: 6px;
            border-radius: 50%;
            background: #00c49f;
            box-shadow:
              0 0 8px rgba(0, 196, 159, 0.7);
          }

          .login-header h1 {
            margin: 0;
            font-size: 30px;
            line-height: 1.2;
            font-weight: 700;
            letter-spacing: -0.5px;
            color: #f4f7fb;
          }

          .login-header p {
            max-width: 330px;
            margin: 10px auto 0;
            font-size: 13px;
            line-height: 1.6;
            color: #7f91a8;
          }

          .login-form {
            margin-top: 32px;
          }

          .login-field {
            margin-bottom: 20px;
          }

          .login-field label {
            display: block;
            margin-bottom: 8px;
            font-size: 12px;
            font-weight: 600;
            letter-spacing: 0.02em;
            color: #a9b7c8;
          }

          .login-password-header {
            display: flex;
            align-items: center;
            justify-content: space-between;
          }

          .login-field input {
            width: 100%;
            height: 48px;
            padding: 0 14px;
            border: 1px solid #263b55;
            border-radius: 9px;
            outline: none;
            background: rgba(4, 13, 24, 0.78);
            color: #e6edf7;
            font-family: inherit;
            font-size: 14px;
            box-sizing: border-box;
            transition:
              border-color 0.2s ease,
              box-shadow 0.2s ease,
              background 0.2s ease;
          }

          .login-field input::placeholder {
            color: #53677f;
          }

          .login-field input:hover {
            border-color: #344c68;
          }

          .login-field input:focus {
            border-color: #00aF91;
            background: rgba(4, 15, 27, 0.92);
            box-shadow:
              0 0 0 3px rgba(0, 196, 159, 0.08);
          }

          .login-button {
            width: 100%;
            height: 48px;
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 10px;
            margin-top: 8px;
            border: 1px solid rgba(0, 196, 159, 0.35);
            border-radius: 9px;
            background:
              linear-gradient(
                135deg,
                #00aF91,
                #00c49f
              );
            color: #03120f;
            font-family: inherit;
            font-size: 14px;
            font-weight: 700;
            cursor: pointer;
            box-shadow:
              0 10px 25px rgba(0, 196, 159, 0.12);
            transition:
              transform 0.2s ease,
              box-shadow 0.2s ease,
              filter 0.2s ease;
          }

          .login-button:hover {
            transform: translateY(-1px);
            filter: brightness(1.05);
            box-shadow:
              0 14px 30px rgba(0, 196, 159, 0.18);
          }

          .login-button:active {
            transform: translateY(0);
          }

          .login-button-arrow {
            font-size: 17px;
            transition: transform 0.2s ease;
          }

          .login-button:hover .login-button-arrow {
            transform: translateX(3px);
          }

          .login-security {
            display: flex;
            gap: 11px;
            margin-top: 24px;
            padding: 14px;
            border: 1px solid #1c3047;
            border-radius: 10px;
            background: rgba(4, 14, 26, 0.55);
          }

          .login-security-icon {
            flex-shrink: 0;
            font-size: 15px;
            line-height: 1.5;
          }

          .login-security strong {
            display: block;
            font-size: 12px;
            font-weight: 650;
            color: #b8c5d5;
          }

          .login-security p {
            margin: 4px 0 0;
            font-size: 11px;
            line-height: 1.5;
            color: #64788f;
          }

          .login-footer {
            display: flex;
            align-items: center;
            justify-content: center;
            flex-wrap: wrap;
            gap: 8px;
            margin-top: 18px;
            font-size: 10px;
            letter-spacing: 0.04em;
            color: #50637a;
            text-align: center;
          }

          .login-footer-separator {
            color: #2f4258;
          }

          @media (max-width: 520px) {
            .login-page {
              padding: 20px 14px;
            }

            .login-card {
              padding: 28px 22px;
              border-radius: 15px;
            }

            .login-header h1 {
              font-size: 27px;
            }

            .login-header p {
              font-size: 12px;
            }

            .login-brand {
              font-size: 17px;
            }
          }
        `}
      </style>
    </div>
  );
}

export default Login;