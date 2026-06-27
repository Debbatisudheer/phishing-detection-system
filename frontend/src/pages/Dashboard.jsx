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

  const [stats, setStats] =
    useState({

      total_analyzed: 0,

      allow: 0,

      suspicious: 0,

      quarantine: 0,

      critical: 0,

    });

  const [health, setHealth] =
    useState({

      database_status: "Healthy",

      auto_cleanup: "Enabled",

      retention: "30 Minutes",

      last_cleanup: "Never",

    });

  useEffect(() => {

    const fetchDashboard =
      async () => {

        try {

          const token =
            localStorage.getItem(
              "token",
            );

          const dashboardResponse =
            await api.get(
              "/api/dashboard",
              {
                headers: {
                  Authorization:
                    `Bearer ${token}`,
                },
              },
            );

          setStats({

            total_analyzed:
              dashboardResponse.data?.total_analyzed || 0,

            allow:
              dashboardResponse.data?.allow || 0,

            suspicious:
              dashboardResponse.data?.suspicious || 0,

            quarantine:
              dashboardResponse.data?.quarantine || 0,

            critical:
              dashboardResponse.data?.critical || 0,

          });

          const healthResponse =
            await api.get(
              "/api/system-health",
              {
                headers: {
                  Authorization:
                    `Bearer ${token}`,
                },
              },
            );

          setHealth(
            healthResponse.data,
          );

        } catch (error) {

          console.error(
            error,
          );

          setStats({

            total_analyzed: 0,

            allow: 0,

            suspicious: 0,

            quarantine: 0,

            critical: 0,

          });

        }

      };

    fetchDashboard();

  }, []);

  const chartData = [

    {

      name: "Allow",

      value: stats.allow,

    },

    {

      name: "Suspicious",

      value: stats.suspicious,

    },

    {

      name: "Quarantine",

      value: stats.quarantine,

    },

  ];

  const COLORS = [

    "#00C49F",

    "#FFBB28",

    "#FF4444",

  ];

  const hasData =
    chartData.some(
      item => item.value > 0,
    );

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

          <h2>Total Analyzed</h2>

          <h1>{stats.total_analyzed}</h1>

        </div>

        <div className="card">

          <h2>Allow</h2>

          <h1>{stats.allow}</h1>

        </div>

        <div className="card">

          <h2>Suspicious</h2>

          <h1>{stats.suspicious}</h1>

        </div>

        <div className="card">

          <h2>Quarantine</h2>

          <h1>{stats.quarantine}</h1>

        </div>

        <div className="card alert-critical">

          <h2>Critical</h2>

          <h1>{stats.critical}</h1>

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

        {

          hasData

          ?

          (

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

                  {

                    chartData.map(

                      (entry,index)=>(

                        <Cell
                          key={index}
                          fill={COLORS[index]}
                        />

                      )

                    )

                  }

                </Pie>

                <Tooltip />

                <Legend />

              </PieChart>

            </ResponsiveContainer>

          )

          :

          (

            <div
  style={{
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    flexDirection: "column",
    height: "300px",
    color: "#666",
    textAlign: "center",
  }}
>

  <h3>

    No Analysis Data Available

  </h3>

  <p>

    Go to the <b>Playground</b> tab and analyze an email or file.

  </p>

  <p>

    The dashboard statistics will appear automatically after the analysis completes.

  </p>

</div>
          )

        }

      </div>

      <div
        className="card"
        style={{
          marginTop: "20px",
        }}
      >

        <h2>

          System Health

        </h2>

        <hr />

        <p>

          <b>

            Database Status

          </b>

          <br />

          🟢 {health.database_status}

        </p>

        <p>

          <b>

            Auto Cleanup

          </b>

          <br />

          {health.auto_cleanup}

        </p>

        <p>

          <b>

            Retention

          </b>

          <br />

          {health.retention}

        </p>

        <p>

          <b>

            Last Cleanup

          </b>

          <br />

          {health.last_cleanup}

        </p>

      </div>

    </div>

  );

}

export default Dashboard;