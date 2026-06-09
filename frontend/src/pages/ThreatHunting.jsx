import {
  PieChart,
  Pie,
  Cell,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";

import { useEffect, useState } from "react";
import api from "../services/api";

function ThreatHunting() {

  const [data, setData] =
    useState(null);

  useEffect(() => {

    loadData();

  }, []);

  const loadData =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/threat-hunting",
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setData(
        response.data,
      );
    };

  if (!data) {

    return <div>Loading...</div>;
  }

  const chartData = [

    {
      name: "Critical",
      value: data.critical_files,
    },

    {
      name: "Quarantine",
      value: data.quarantine_files,
    },
  ];

  const COLORS = [
    "#FF4444",
    "#FF8042",
  ];

  return (

    <div
      style={{
        padding: "20px",
      }}
    >

      <h1>
        Threat Hunting Dashboard
      </h1>

      <div
        style={{
          display: "flex",
          gap: "20px",
          marginBottom: "30px",
        }}
      >

        <div
          style={{
            border:
              "1px solid #333",
            padding: "20px",
            flex: 1,
          }}
        >
          <h3>
            Critical Files
          </h3>

          <h2>
            {
              data.critical_files
            }
          </h2>
        </div>

        <div
          style={{
            border:
              "1px solid #333",
            padding: "20px",
            flex: 1,
          }}
        >
          <h3>
            Quarantine Files
          </h3>

          <h2>
            {
              data.quarantine_files
            }
          </h2>
        </div>

      </div>

      <ResponsiveContainer
        width="100%"
        height={400}
      >

        <PieChart>

          <Pie
            data={chartData}
            dataKey="value"
            nameKey="name"
            outerRadius={150}
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
                    COLORS[
                      index %
                      COLORS.length
                    ]
                  }
                />
              ),
            )}

          </Pie>

          <Tooltip />

          <Legend />

        </PieChart>

      </ResponsiveContainer>

      <h2>
        MITRE Techniques
      </h2>

      <table
        border="1"
        width="100%"
      >

        <thead>

          <tr>

            <th>
              Technique
            </th>

            <th>
              Count
            </th>

          </tr>

        </thead>

        <tbody>

          {Object.entries(
            data.top_mitre,
          ).map(
            (
              [technique, count],
              index,
            ) => (

              <tr key={index}>

                <td>
                  {technique}
                </td>

                <td>
                  {count}
                </td>

              </tr>
            ),
          )}

        </tbody>

      </table>

    </div>
  );
}

export default ThreatHunting;