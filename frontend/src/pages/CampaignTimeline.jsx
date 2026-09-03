import { useEffect, useState } from "react";
import api from "../services/api";
import "./CampaignTimeline.css";

function CampaignTimeline() {
  const [timeline, setTimeline] =
    useState([]);

  useEffect(() => {
    loadTimeline();
  }, []);

  const loadTimeline =
    async () => {
      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/campaign-timeline",
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setTimeline(
        response.data,
      );
    };

  return (
    <main className="campaign-timeline-page">

      <div className="campaign-timeline-container">

        {/* =================================================
            HEADER
            ================================================= */}

        <header className="campaign-timeline-header">

          <div className="campaign-timeline-title-area">

            <div className="campaign-timeline-label">
              <span className="campaign-timeline-dot" />
              Campaign Intelligence
            </div>

            <h1>
              Campaign Timeline
            </h1>

          </div>

        </header>


        {/* =================================================
            TIMELINE PANEL
            ================================================= */}

        <section className="campaign-timeline-panel">

          <div className="campaign-timeline-panel-header">

            <h2>
              Campaign Timeline
            </h2>

          </div>


          <div className="campaign-timeline-table-wrapper">

            <table className="campaign-timeline-table">

              <thead>

                <tr>

                  <th>
                    IOC
                  </th>

                  <th>
                    First Seen
                  </th>

                  <th>
                    Last Seen
                  </th>

                  <th>
                    Occurrences
                  </th>

                </tr>

              </thead>


              <tbody>

                {timeline.map(
                  (
                    item,
                    index,
                  ) => (

                    <tr
                      key={index}
                    >

                      <td>

                        <div className="timeline-ioc">

                          <span className="ioc-indicator" />

                          <span className="ioc-text">
                            {item.ioc}
                          </span>

                        </div>

                      </td>


                      <td>

                        <span className="timeline-date">
                          {item.first_seen}
                        </span>

                      </td>


                      <td>

                        <span className="timeline-date">
                          {item.last_seen}
                        </span>

                      </td>


                      <td>

                        <span className="occurrences-badge">
                          {item.occurrences}
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

export default CampaignTimeline;