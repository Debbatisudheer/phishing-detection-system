function ErrorCard({

  icon,
  title,
  message,
  buttonText,
  onClick,

}) {

  return (

    <div
      style={{

        minHeight: "80vh",

        display: "flex",

        justifyContent: "center",

        alignItems: "center",

        background: "#0f172a",

        padding: "30px",

      }}
    >

      <div
        style={{

          width: "100%",

          maxWidth: "560px",

          background: "#1e293b",

          border: "1px solid #ef4444",

          borderRadius: "16px",

          padding: "35px",

          textAlign: "center",

          boxShadow:
            "0 15px 40px rgba(0,0,0,.45)",

        }}
      >

        <div
          style={{
            fontSize: "70px",
            marginBottom: "15px",
          }}
        >
          {icon}
        </div>

        <h1
          style={{

            color: "#ffffff",

            marginBottom: "15px",

          }}
        >
          {title}
        </h1>

        <p
          style={{

            color: "#cbd5e1",

            lineHeight: "28px",

            marginBottom: "30px",

          }}
        >
          {message}
        </p>

        {

          buttonText && (

            <button

              onClick={onClick}

              style={{

                background: "#2563eb",

                color: "white",

                border: "none",

                padding: "12px 30px",

                borderRadius: "8px",

                cursor: "pointer",

                fontSize: "16px",

              }}

            >

              {buttonText}

            </button>

          )

        }

      </div>

    </div>

  );

}

export default ErrorCard;