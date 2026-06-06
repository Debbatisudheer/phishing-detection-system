import { useEffect, useState } from "react";
import api from "../services/api";

function ThreatIntel() {

  const [findings, setFindings] =
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
              "/api/recent-findings",
              {
                headers: {
                  Authorization:
                    `Bearer ${token}`,
                },
              },
            );

          setFindings(
            response.data,
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
        Top Risk Files
      </h2>

      {findings
        .filter(
          (
            item,
          ) =>
            item.risk_level ===
              "HIGH" ||
            item.risk_level ===
              "CRITICAL",
        )
        .map(
          (
            item,
            index,
          ) => (

            <div
              key={index}
              className="card alert-critical"
            >

              <h3>
                {item.file_name}
              </h3>

              <p>
                Risk:
                {" "}
                {item.risk_level}
              </p>

              <p>
                Score:
                {" "}
                {item.risk_score}
              </p>

              <p>
                Verdict:
                {" "}
                {item.verdict}
              </p>

            </div>
          ),
        )}

    </div>
  );
}

export default ThreatIntel;