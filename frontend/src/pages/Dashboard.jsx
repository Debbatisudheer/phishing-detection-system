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
  const [stats, setStats] = useState({
    total_analyzed: 0,
    allow: 0,
    suspicious: 0,
    quarantine: 0,
    critical: 0,
  });

  const [health, setHealth] = useState({
    database_status: "Healthy",
    auto_cleanup: "Enabled",
    retention: "30 Minutes",
    last_cleanup: "Never",
  });

  useEffect(() => {
    const fetchDashboard = async () => {
      try {
        const token = localStorage.getItem("token");

        const dashboardResponse = await api.get("/api/dashboard", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

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

        const healthResponse = await api.get(
          "/api/system-health",
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );

        setHealth(healthResponse.data);
      } catch (error) {
        console.error(error);

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

  const hasData = chartData.some(
    (item) => item.value > 0
  );

  return (
    <div
      style={{
        minHeight: "100vh",
        padding: "32px",
        background:
          "linear-gradient(135deg, #07111f 0%, #0b1628 50%, #08111f 100%)",
        color: "#e6edf7",
        boxSizing: "border-box",
      }}
    >
      {/* ============================================================
          Dashboard Header
          ============================================================ */}
      <div
        style={{
          marginBottom: "28px",
          display: "flex",
          flexDirection: "column",
          gap: "6px",
        }}
      >
        <h1
          style={{
            margin: 0,
            fontSize: "30px",
            fontWeight: 700,
            letterSpacing: "-0.5px",
            color: "#f4f7fb",
          }}
        >
          SOC Dashboard
        </h1>

        <p
          style={{
            margin: 0,
            fontSize: "14px",
            color: "#8fa1b8",
          }}
        >
          Security operations overview and analysis status
        </p>
      </div>

      {/* ============================================================
          Statistics
          ============================================================ */}
      <div className="dashboard-grid">
        {/* Total Analyzed */}
        <div className="dashboard-stat-card">
          <div className="dashboard-stat-label">
            Total Analyzed
          </div>

          <div className="dashboard-stat-value">
            {stats.total_analyzed}
          </div>
        </div>

        {/* Allow */}
        <div className="dashboard-stat-card">
          <div className="dashboard-stat-label">
            Allow
          </div>

          <div
            className="dashboard-stat-value"
            style={{
              color: "#00C49F",
            }}
          >
            {stats.allow}
          </div>
        </div>

        {/* Suspicious */}
        <div className="dashboard-stat-card">
          <div className="dashboard-stat-label">
            Suspicious
          </div>

          <div
            className="dashboard-stat-value"
            style={{
              color: "#FFBB28",
            }}
          >
            {stats.suspicious}
          </div>
        </div>

        {/* Quarantine */}
        <div className="dashboard-stat-card">
          <div className="dashboard-stat-label">
            Quarantine
          </div>

          <div
            className="dashboard-stat-value"
            style={{
              color: "#FF4444",
            }}
          >
            {stats.quarantine}
          </div>
        </div>

        {/* Critical */}
        <div className="dashboard-stat-card dashboard-critical-card">
          <div className="dashboard-stat-label">
            Critical
          </div>

          <div
            className="dashboard-stat-value"
            style={{
              color: "#ff5c5c",
            }}
          >
            {stats.critical}
          </div>
        </div>
      </div>

      {/* ============================================================
          Risk Distribution
          ============================================================ */}
      <div className="dashboard-card dashboard-chart-card">
        <div className="dashboard-section-header">
          <div>
            <h2 className="dashboard-section-title">
              Risk Distribution
            </h2>

            <p className="dashboard-section-description">
              Distribution of analyzed items by security classification
            </p>
          </div>
        </div>

        <div className="dashboard-chart-container">
          {hasData ? (
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
                    (entry, index) => (
                      <Cell
                        key={index}
                        fill={COLORS[index]}
                      />
                    )
                  )}
                </Pie>

                <Tooltip />

                <Legend />
              </PieChart>
            </ResponsiveContainer>
          ) : (
            <div className="dashboard-empty-state">
              <div className="dashboard-empty-icon">
                —
              </div>

              <h3>
                No Analysis Data Available
              </h3>

              <p>
                Go to the <b>Playground</b> tab and
                analyze an email or file.
              </p>

              <p>
                The dashboard statistics will appear
                automatically after the analysis completes.
              </p>
            </div>
          )}
        </div>
      </div>

      {/* ============================================================
          System Health
          ============================================================ */}
      <div className="dashboard-card dashboard-health-card">
        <div className="dashboard-section-header">
          <div>
            <h2 className="dashboard-section-title">
              System Health
            </h2>

            <p className="dashboard-section-description">
              Current status of core platform services
            </p>
          </div>
        </div>

        <div className="dashboard-health-divider" />

        <div className="dashboard-health-grid">
          {/* Database Status */}
          <div className="dashboard-health-item">
            <span className="dashboard-health-label">
              Database Status
            </span>

            <span className="dashboard-health-value">
              <span
                style={{
                  color: "#00C49F",
                  marginRight: "8px",
                }}
              >
                ●
              </span>

              {health.database_status}
            </span>
          </div>

          {/* Auto Cleanup */}
          <div className="dashboard-health-item">
            <span className="dashboard-health-label">
              Auto Cleanup
            </span>

            <span className="dashboard-health-value">
              {health.auto_cleanup}
            </span>
          </div>

          {/* Retention */}
          <div className="dashboard-health-item">
            <span className="dashboard-health-label">
              Retention
            </span>

            <span className="dashboard-health-value">
              {health.retention}
            </span>
          </div>

          {/* Last Cleanup */}
          <div className="dashboard-health-item">
            <span className="dashboard-health-label">
              Last Cleanup
            </span>

            <span className="dashboard-health-value">
              {health.last_cleanup}
            </span>
          </div>
        </div>
      </div>

      {/* ============================================================
          Dashboard Styles
          ============================================================ */}
      <style>
        {`
          .dashboard-grid {
            display: grid;
            grid-template-columns: repeat(5, minmax(0, 1fr));
            gap: 16px;
          }

          .dashboard-stat-card {
            position: relative;
            overflow: hidden;
            min-height: 130px;
            padding: 22px;
            border: 1px solid #1d3049;
            border-radius: 14px;
            background:
              linear-gradient(
                145deg,
                rgba(15, 31, 52, 0.96),
                rgba(9, 22, 39, 0.96)
              );
            box-shadow:
              0 12px 30px rgba(0, 0, 0, 0.18);
            transition:
              transform 0.2s ease,
              border-color 0.2s ease,
              box-shadow 0.2s ease;
            box-sizing: border-box;
          }

          .dashboard-stat-card::before {
            content: "";
            position: absolute;
            top: 0;
            left: 0;
            right: 0;
            height: 2px;
            background: rgba(100, 255, 218, 0.45);
          }

          .dashboard-stat-card:hover {
            transform: translateY(-2px);
            border-color: #2b4565;
            box-shadow:
              0 16px 36px rgba(0, 0, 0, 0.25);
          }

          .dashboard-critical-card {
            border-color: rgba(255, 68, 68, 0.28);
          }

          .dashboard-critical-card::before {
            background: rgba(255, 68, 68, 0.65);
          }

          .dashboard-stat-label {
            font-size: 13px;
            font-weight: 600;
            letter-spacing: 0.02em;
            color: #8fa1b8;
          }

          .dashboard-stat-value {
            margin-top: 18px;
            font-size: 34px;
            line-height: 1;
            font-weight: 700;
            color: #f4f7fb;
          }

          .dashboard-card {
            border: 1px solid #1d3049;
            border-radius: 14px;
            background:
              linear-gradient(
                145deg,
                rgba(15, 31, 52, 0.96),
                rgba(9, 22, 39, 0.96)
              );
            box-shadow:
              0 12px 30px rgba(0, 0, 0, 0.18);
          }

          .dashboard-chart-card {
            margin-top: 20px;
            min-height: 460px;
            padding: 24px;
            box-sizing: border-box;
          }

          .dashboard-health-card {
            margin-top: 20px;
            padding: 24px;
            box-sizing: border-box;
          }

          .dashboard-section-header {
            display: flex;
            align-items: center;
            justify-content: space-between;
            gap: 20px;
          }

          .dashboard-section-title {
            margin: 0;
            font-size: 19px;
            font-weight: 650;
            color: #f4f7fb;
          }

          .dashboard-section-description {
            margin: 6px 0 0;
            font-size: 13px;
            line-height: 1.5;
            color: #71839b;
          }

          .dashboard-chart-container {
            width: 100%;
            height: 370px;
            margin-top: 8px;
          }

          .dashboard-empty-state {
            height: 300px;
            display: flex;
            align-items: center;
            justify-content: center;
            flex-direction: column;
            text-align: center;
            padding: 20px;
            box-sizing: border-box;
            color: #71839b;
          }

          .dashboard-empty-icon {
            width: 42px;
            height: 42px;
            display: flex;
            align-items: center;
            justify-content: center;
            margin-bottom: 14px;
            border: 1px solid #2a3d56;
            border-radius: 50%;
            color: #71839b;
            font-size: 20px;
          }

          .dashboard-empty-state h3 {
            margin: 0;
            font-size: 16px;
            font-weight: 600;
            color: #b8c5d6;
          }

          .dashboard-empty-state p {
            max-width: 560px;
            margin: 8px 0 0;
            font-size: 13px;
            line-height: 1.6;
            color: #71839b;
          }

          .dashboard-health-divider {
            height: 1px;
            margin: 20px 0;
            background: #1d3049;
          }

          .dashboard-health-grid {
            display: grid;
            grid-template-columns: repeat(4, minmax(0, 1fr));
            gap: 14px;
          }

          .dashboard-health-item {
            min-height: 90px;
            padding: 17px;
            border: 1px solid #1b2d45;
            border-radius: 10px;
            background: rgba(5, 16, 30, 0.48);
            box-sizing: border-box;
          }

          .dashboard-health-label {
            display: block;
            font-size: 12px;
            font-weight: 600;
            color: #71839b;
          }

          .dashboard-health-value {
            display: block;
            margin-top: 12px;
            font-size: 14px;
            font-weight: 600;
            color: #d7e0eb;
          }

          @media (max-width: 1100px) {
            .dashboard-grid {
              grid-template-columns: repeat(3, minmax(0, 1fr));
            }

            .dashboard-health-grid {
              grid-template-columns: repeat(2, minmax(0, 1fr));
            }
          }

          @media (max-width: 700px) {
            .dashboard-grid {
              grid-template-columns: repeat(2, minmax(0, 1fr));
            }

            .dashboard-chart-card,
            .dashboard-health-card {
              padding: 18px;
            }

            .dashboard-chart-container {
              height: 330px;
            }
          }

          @media (max-width: 520px) {
            .dashboard-grid {
              grid-template-columns: 1fr;
            }

            .dashboard-health-grid {
              grid-template-columns: 1fr;
            }

            .dashboard-chart-card {
              min-height: 430px;
            }

            .dashboard-chart-container {
              height: 320px;
            }
          }
        `}
      </style>
    </div>
  );
}

export default Dashboard;