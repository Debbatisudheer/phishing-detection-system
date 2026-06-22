import { useEffect, useState } from "react";
import api from "../services/api";

function ThreatIntel() {

  const [stats, setStats] =
    useState({});

  const [files, setFiles] =
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
              "/api/threat-intel",
              {
                headers: {
                  Authorization:
                    `Bearer ${token}`,
                },
              },
            );

          setStats(
            response.data.stats,
          );

          setFiles(
            response.data.top_files,
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
        Threat Intelligence Dashboard
      </h1>

      <h2>
        Statistics
      </h2>

      <div
        className="card"
      >
        <p>
          Total IOCs:
          {" "}
          {stats.total_iocs}
        </p>

        <p>
          Total Alerts:
          {" "}
          {stats.total_alerts}
        </p>

        <p>
          Critical Files:
          {" "}
          {stats.critical_files}
        </p>
      </div>

      <h2>
        Top Risk Files
      </h2>

      {files.map(
        (
          item,
          index,
        ) => (

          <div
            key={index}
            className="card alert-critical"
          >

            <h3>
              {item.file}
            </h3>

            <p>
              Risk:
              {" "}
              {item.level}
            </p>

            <p>
              Score:
              {" "}
              {item.score}
            </p>

          </div>
        ),
      )}

    </div>
  );
}

export default ThreatIntel;