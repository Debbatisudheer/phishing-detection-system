import { useEffect, useState } from "react";
import api from "../services/api";

function MITREDashboard() {

  const [stats, setStats] =
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
              "/api/mitre-stats",
              {
                headers: {
                  Authorization:
                    `Bearer ${token}`,
                },
              },
            );

          setStats(
            Array.isArray(
              response.data,
            )
              ? response.data
              : [],
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
        MITRE Dashboard
      </h1>

      {stats.map(
        (
          item,
          index,
        ) => (

          <div
            key={index}
            className="card"
          >

            <h3>
              Techniques
            </h3>

            <pre>
              {item.technique}
            </pre>

            <p>
              Count:
              {" "}
              {item.count}
            </p>

          </div>
        ),
      )}

    </div>
  );
}

export default MITREDashboard;