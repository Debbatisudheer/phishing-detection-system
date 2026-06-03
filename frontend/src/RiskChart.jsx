import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  CartesianGrid,
  ResponsiveContainer,
} from "recharts";

function RiskChart({ emails }) {

  const chartData = emails.map(
    (email, index) => ({

      name: `Alert ${index + 1}`,

      risk: email.risk_score,
    })
  );

  return (

    <div
      style={{
        backgroundColor: "#1e1e1e",
        padding: "20px",
        borderRadius: "10px",
        marginBottom: "20px",
      }}
    >

      <h2>
        Risk Score Trends
      </h2>

      <ResponsiveContainer
        width="100%"
        height={300}
      >

        <LineChart data={chartData}>

          <CartesianGrid strokeDasharray="3 3" />

          <XAxis dataKey="name" />

          <YAxis />

          <Tooltip />

          <Line
            type="monotone"
            dataKey="risk"
            stroke="#ff4d4f"
            strokeWidth={3}
          />

        </LineChart>

      </ResponsiveContainer>

    </div>
  );
}

export default RiskChart;