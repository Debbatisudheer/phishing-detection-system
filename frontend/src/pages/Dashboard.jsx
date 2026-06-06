import { useEffect, useState } from "react";
import api from "../services/api";

import {
  PieChart,
  Pie,
  Cell,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";

function Dashboard() {

  const [stats, setStats] = useState({});

  useEffect(() => {

    const fetchDashboard = async () => {

      try {

        const token =
          localStorage.getItem(
            "token",
          );

        const response =
          await api.get(
            "/api/dashboard",
            {
              headers: {
                Authorization:
                  `Bearer ${token}`,
              },
            },
          );

        setStats(
          response.data,
        );

      } catch (error) {

        console.error(
          error,
        );
      }
    };

    fetchDashboard();

  }, []);

  const chartData = [

    {
      name: "Allow",
      value: stats.allow || 0,
    },

    {
      name: "Suspicious",
      value: stats.suspicious || 0,
    },

    {
      name: "Quarantine",
      value: stats.quarantine || 0,
    },
  ];

  const COLORS = [
    "#00C49F",
    "#FFBB28",
    "#FF4444",
  ];

  return (

    <div
      style={{
        padding: "20px",
      }}
    >

      <h1>
        SOC Dashboard
      </h1>

      <div className="dashboard-grid">

        <div className="card">

          <h2>
            Total Analyzed
          </h2>

          <h1>
            {stats.total_analyzed || 0}
          </h1>

        </div>

        <div className="card">

          <h2>
            Allow
          </h2>

          <h1>
            {stats.allow || 0}
          </h1>

        </div>

        <div className="card">

          <h2>
            Suspicious
          </h2>

          <h1>
            {stats.suspicious || 0}
          </h1>

        </div>

        <div className="card">

          <h2>
            Quarantine
          </h2>

          <h1>
            {stats.quarantine || 0}
          </h1>

        </div>

        <div className="card alert-critical">

          <h2>
            Critical
          </h2>

          <h1>
            {stats.critical || 0}
          </h1>

        </div>

      </div>

      <div
        className="card"
        style={{
          marginTop: "20px",
          height: "400px",
        }}
      >

        <h2>
          Risk Distribution
        </h2>

        <ResponsiveContainer
          width="100%"
          height="100%"
        >

          <PieChart>

            <Pie
              data={chartData}
              cx="50%"
              cy="50%"
              outerRadius={120}
              dataKey="value"
              label
            >

              {chartData.map(
                (
                  entry,
                  index,
                ) => (

                  <Cell
                    key={index}
                    fill={
                      COLORS[index]
                    }
                  />

                ),
              )}

            </Pie>

            <Tooltip />

            <Legend />

          </PieChart>

        </ResponsiveContainer>

      </div>

    </div>
  );
}

export default Dashboard;