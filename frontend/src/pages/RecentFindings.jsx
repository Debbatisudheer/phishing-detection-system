import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import api from "../services/api";
import "./RecentFindings.css";

function RecentFindings() {
  const [findings, setFindings] =
    useState([]);

  useEffect(() => {
    const fetchFindings =
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

    fetchFindings();
  }, []);

  return (
    <main className="recent-findings-page">

      <div className="recent-findings-container">

        {/* =================================================
            HEADER
            ================================================= */}

        <header className="recent-findings-header">

          <div className="recent-findings-title-area">

            <div className="recent-findings-status">
              <span className="status-dot" />
              Recent Activity
            </div>

            <h1>
              Recent Findings
            </h1>

          </div>

        </header>


        {/* =================================================
            FINDINGS
            ================================================= */}

        <section className="findings-section">

          {findings.map(
            (
              finding,
              index,
            ) => (

              <article
                key={index}
                className="finding-card"
              >

                {/* Card Header */}

                <div className="finding-card-header">

                  <div className="finding-file-info">

                    <div className="file-icon">
                      <span />
                    </div>

                    <div className="file-name-wrapper">

                      <Link
                        to={`/file/${finding.file_name}`}
                        className="finding-file-name"
                      >
                        {finding.file_name}
                      </Link>

                    </div>

                  </div>

                  <div className="finding-index">
                    #{String(
                      index + 1,
                    ).padStart(2, "0")}
                  </div>

                </div>


                {/* Divider */}

                <div className="finding-divider" />


                {/* Finding Details */}

                <div className="finding-details">

                  <div className="finding-detail">

                    <span className="detail-label">
                      Risk Score
                    </span>

                    <span className="detail-value risk-score">
                      {finding.risk_score}
                    </span>

                  </div>


                  <div className="finding-detail">

                    <span className="detail-label">
                      Risk Level
                    </span>

                    <span
                      className={`detail-value risk-level risk-${String(
                        finding.risk_level ||
                          "",
                      ).toLowerCase()}`}
                    >
                      {finding.risk_level}
                    </span>

                  </div>


                  <div className="finding-detail">

                    <span className="detail-label">
                      Verdict
                    </span>

                    <span
                      className={`detail-value verdict verdict-${String(
                        finding.verdict ||
                          "",
                      ).toLowerCase()}`}
                    >
                      {finding.verdict}
                    </span>

                  </div>

                </div>

              </article>

            ),
          )}

        </section>

      </div>

    </main>
  );
}

export default RecentFindings;