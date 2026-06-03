function DashboardCards({ emails }) {

  const totalAlerts = emails.length;

  const highRiskAlerts =
    emails.filter(
      (email) =>
        email.risk_score >= 200
    ).length;

  const quarantined =
    emails.filter(
      (email) =>
        email.decision ===
        "QUARANTINE"
    ).length;

  return (

    <div
      style={{
        display: "flex",
        gap: "20px",
        marginBottom: "20px",
      }}
    >

      <Card
        title="Total Alerts"
        value={totalAlerts}
      />

      <Card
        title="High Risk"
        value={highRiskAlerts}
      />

      <Card
        title="Quarantined"
        value={quarantined}
      />

    </div>
  );
}

function Card({ title, value }) {

  return (

    <div
      style={{
        backgroundColor: "#1e1e1e",
        padding: "20px",
        borderRadius: "10px",
        width: "220px",
        boxShadow:
          "0px 0px 10px rgba(0,0,0,0.5)",
      }}
    >

      <h3>{title}</h3>

      <h1>{value}</h1>

    </div>
  );
}

export default DashboardCards;