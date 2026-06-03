function ThreatIntelPanel({ emails }) {

  return (

    <div
      style={{
        backgroundColor: "#1e1e1e",
        padding: "20px",
        borderRadius: "10px",
        marginTop: "20px",
      }}
    >

      <h2>
        Threat Intelligence Panel
      </h2>

      {emails.map((email, index) => (

        <div
          key={index}
          style={{
            border:
              "1px solid #333",
            padding: "10px",
            marginBottom: "10px",
            borderRadius: "5px",
          }}
        >

          <p>
            <strong>
              Subject:
            </strong>
            {" "}
            {email.subject}
          </p>

          <p>
            <strong>
              Findings:
            </strong>
          </p>

          <p
            style={{
              color: "#ff4d4f",
            }}
          >
            {email.findings}
          </p>

        </div>

      ))}

    </div>
  );
}

export default ThreatIntelPanel;