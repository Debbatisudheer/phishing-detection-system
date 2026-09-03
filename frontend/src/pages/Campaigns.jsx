import { useEffect, useState } from "react";
import api from "../services/api";
import "./Campaigns.css";

function Campaigns() {
  const [campaigns, setCampaigns] =
    useState([]);

  const [sources, setSources] =
    useState([]);

  useEffect(() => {
    loadCampaigns();
  }, []);

  const loadCampaigns =
    async () => {
      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/campaigns",
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setCampaigns(
        response.data,
      );
    };

  const loadSources =
    async (ioc) => {
      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          `/api/ioc-sources?ioc=${encodeURIComponent(ioc)}`,
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setSources(
        response.data,
      );
    };

  return (
    <main className="campaigns-page">

      <div className="campaigns-container">

        {/* =================================================
            HEADER
            ================================================= */}

        <header className="campaigns-header">

          <div className="campaigns-title-area">

            <div className="campaigns-section-label">
              <span className="campaigns-live-dot" />
              Threat Campaign Intelligence
            </div>

            <h1>
              Campaign Dashboard
            </h1>

          </div>

        </header>


        {/* =================================================
            CAMPAIGNS PANEL
            ================================================= */}

        <section className="campaign-panel">

          <div className="campaign-panel-header">

            <div>
              <h2>
                Campaign Dashboard
              </h2>
            </div>

          </div>


          <div className="campaign-table-wrapper">

            <table className="campaign-table">

              <thead>
                <tr>

                  <th>
                    IOC
                  </th>

                  <th>
                    Count
                  </th>

                  <th>
                    Severity
                  </th>

                  <th>
                    Action
                  </th>

                </tr>
              </thead>


              <tbody>

                {campaigns.map(
                  (
                    campaign,
                    index,
                  ) => (

                    <tr
                      key={index}
                    >

                      <td>
                        <span className="ioc-value">
                          {campaign.ioc}
                        </span>
                      </td>

                      <td>
                        <span className="count-value">
                          {campaign.count}
                        </span>
                      </td>

                      <td>
                        <span
                          className={`severity-badge severity-${String(
                            campaign.severity ||
                              "",
                          ).toLowerCase()}`}
                        >
                          <span className="severity-dot" />

                          {campaign.severity}
                        </span>
                      </td>

                      <td>

                        <button
                          className="sources-button"
                          onClick={() =>
                            loadSources(
                              campaign.ioc,
                            )
                          }
                        >
                          View Sources
                        </button>

                      </td>

                    </tr>

                  ),
                )}

              </tbody>

            </table>

          </div>

        </section>


        {/* =================================================
            IOC SOURCES
            ================================================= */}

        <section className="sources-panel">

          <div className="sources-panel-header">

            <div>

              <h2>
                IOC Sources
              </h2>

            </div>

          </div>


          <div className="sources-table-wrapper">

            <table className="sources-table">

              <thead>

                <tr>

                  <th>
                    Source
                  </th>

                  <th>
                    File
                  </th>

                  <th>
                    Time
                  </th>

                </tr>

              </thead>


              <tbody>

                {sources.map(
                  (
                    source,
                    index,
                  ) => (

                    <tr
                      key={index}
                    >

                      <td>
                        <span className="source-type">
                          {source.source_type}
                        </span>
                      </td>

                      <td>
                        <span className="source-file">
                          {source.file_name}
                        </span>
                      </td>

                      <td>
                        <span className="source-time">
                          {source.created_at}
                        </span>
                      </td>

                    </tr>

                  ),
                )}

              </tbody>

            </table>

          </div>

        </section>

      </div>

    </main>
  );
}

export default Campaigns;