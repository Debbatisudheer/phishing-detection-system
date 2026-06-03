function AttackTimeline({ emails }) {

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
        Attack Timeline
      </h2>

      {emails.map((email, index) => (

        <div
          key={index}
          style={{
            borderLeft:
              "4px solid red",
            paddingLeft: "10px",
            marginBottom: "15px",
          }}
        >

          <p>
            <strong>
              {email.subject}
            </strong>
          </p>

          <p>
            Sender:
            {" "}
            {email.sender}
          </p>

          <p>
            Risk Score:
            {" "}
            {email.risk_score}
          </p>

          <p>
            Decision:
            {" "}
            {email.decision}
          </p>

        </div>

      ))}

    </div>
  );
}

export default AttackTimeline;