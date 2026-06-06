import { useEffect, useState } from "react";
import api from "../services/api";

function IncidentDashboard() {

  const [stats, setStats] =
    useState({});

  const [incidents, setIncidents] =
    useState([]);

  useEffect(() => {

    const fetchData =
      async () => {

        try {

          const token =
            localStorage.getItem(
              "token",
            );

          const response =
            await api.get(
              "/api/incident-dashboard",
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

          const recentResponse =
            await api.get(
              "/api/recent-incidents",
              {
                headers: {
                  Authorization:
                    `Bearer ${token}`,
                },
              },
            );

          setIncidents(
            recentResponse.data,
          );

        } catch (error) {

          console.error(
            error,
          );
        }
      };

    fetchData();

  }, []);

  return (

    <div
      style={{
        padding: "20px",
      }}
    >

      <h1>
        Incident Dashboard
      </h1>

      <div
        className="dashboard-grid"
      >

        <div className="card">

          <h2>
            Total Incidents
          </h2>

          <h1>
            {stats.total || 0}
          </h1>

        </div>

        <div className="card">

          <h2>
            Open Incidents
          </h2>

          <h1>
            {stats.open || 0}
          </h1>

        </div>

        <div className="card">

          <h2>
            Closed Incidents
          </h2>

          <h1>
            {stats.closed || 0}
          </h1>

        </div>

      </div>

      <hr />

      <h2>
        Recent Incidents
      </h2>

      <table
        border="1"
        cellPadding="10"
        style={{
          width: "100%",
          marginTop: "20px",
        }}
      >

        <thead>

          <tr>

            <th>ID</th>

            <th>File</th>

            <th>Analyst</th>

            <th>Status</th>

          </tr>

        </thead>

        <tbody>

          {incidents.map(
            (
              incident,
              index,
            ) => (

              <tr
                key={index}
              >

                <td>
                  {incident.id}
                </td>

                <td>
                  {incident.file_name}
                </td>

                <td>
                  {incident.analyst}
                </td>

                <td>
                  {incident.status}
                </td>

              </tr>
            ),
          )}

        </tbody>

      </table>

    </div>
  );
}

export default IncidentDashboard;