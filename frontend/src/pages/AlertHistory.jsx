import { useEffect, useState } from "react";
import api from "../services/api";

function AlertHistory() {

  const [alerts, setAlerts] =
    useState([]);

  useEffect(() => {

    loadAlerts();

  }, []);

  const loadAlerts =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/alerts",
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setAlerts(
        response.data,
      );
    };

  return (

    <div>

      <h1>
        Alert History
      </h1>

      <table
        border="1"
        width="100%"
      >

        <thead>

          <tr>

            <th>ID</th>

            <th>Time</th>

            <th>File</th>

            <th>Risk</th>

            <th>Verdict</th>

          </tr>

        </thead>

        <tbody>

          {alerts.map(
            (alert) => (

              <tr
                key={alert.id}
              >

                <td>
                  {alert.id}
                </td>

                <td>
                  {alert.alert_time}
                </td>

                <td>
                  {alert.file_name}
                </td>

                <td>
                  {alert.risk_level}
                </td>

                <td>
                  {alert.verdict}
                </td>

              </tr>
            ),
          )}

        </tbody>

      </table>

    </div>
  );
}

export default AlertHistory;