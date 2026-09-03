import { useState } from "react";
import api from "../services/api";
import "./UpdateCase.css";

function UpdateCase() {
  const [caseId, setCaseId] =
    useState("");

  const [status, setStatus] =
    useState("");

  const [notes, setNotes] =
    useState("");

  const [timelineNote, setTimelineNote] =
    useState("");

  const [caseNotes, setCaseNotes] =
    useState([]);

  const updateCase =
    async () => {
      const token =
        localStorage.getItem(
          "token",
        );

      await api.put(
        `/api/case-details/${caseId}`,
        {
          status,
          notes,
        },
        {
          headers: {
            Authorization:
              `Bearer ${token}`,
          },
        },
      );

      alert(
        "Case Updated",
      );
    };

  const addNote =
    async () => {
      const token =
        localStorage.getItem(
          "token",
        );

      await api.post(
        "/api/case-note",
        {
          case_id:
            parseInt(
              caseId,
            ),
          analyst:
            "Sudheer",
          note:
            timelineNote,
        },
        {
          headers: {
            Authorization:
              `Bearer ${token}`,
          },
        },
      );

      alert(
        "Note Added",
      );

      setTimelineNote(
        "",
      );

      loadNotes();
    };

  const loadNotes =
    async () => {
      if (!caseId) {
        return;
      }

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          `/api/case-notes/${caseId}`,
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setCaseNotes(
        response.data,
      );
    };

  return (
    <main className="update-case-page">

      <div className="update-case-container">

        {/* =================================================
            HEADER
            ================================================= */}

        <header className="update-case-header">

          <div className="update-case-heading">

            <div className="update-case-label">
              <span className="case-status-dot" />
              Case Management
            </div>

            <h1>
              Update Case
            </h1>

          </div>

        </header>


        {/* =================================================
            CASE WORKSPACE
            ================================================= */}

        <section className="case-workspace">

          {/* =================================================
              CASE DETAILS
              ================================================= */}

          <div className="case-panel">

            <div className="case-panel-header">

              <div>
                <h2>
                  Case Details
                </h2>
              </div>

            </div>


            <div className="case-panel-body">

              {/* Case ID */}

              <div className="case-field">

                <label>
                  Case ID
                </label>

                <input
                  placeholder="Case ID"
                  value={caseId}
                  onChange={(e) =>
                    setCaseId(
                      e.target.value,
                    )
                  }
                />

              </div>


              <button
                className="secondary-action"
                onClick={loadNotes}
              >
                Load Notes
              </button>


              {/* Status */}

              <div className="case-field">

                <label>
                  Status
                </label>

                <input
                  placeholder="Status"
                  value={status}
                  onChange={(e) =>
                    setStatus(
                      e.target.value,
                    )
                  }
                />

              </div>


              {/* Case Notes */}

              <div className="case-field">

                <label>
                  Case Notes
                </label>

                <textarea
                  placeholder="Case Notes"
                  value={notes}
                  onChange={(e) =>
                    setNotes(
                      e.target.value,
                    )
                  }
                />

              </div>


              <button
                className="primary-action"
                onClick={updateCase}
              >
                Update Case
              </button>

            </div>

          </div>


          {/* =================================================
              ANALYST TIMELINE
              ================================================= */}

          <div className="case-panel timeline-panel">

            <div className="case-panel-header">

              <div>
                <h2>
                  Analyst Timeline
                </h2>
              </div>

            </div>


            <div className="case-panel-body">

              <div className="case-field">

                <label>
                  Add Investigation Note
                </label>

                <textarea
                  placeholder="Add Investigation Note"
                  value={timelineNote}
                  onChange={(e) =>
                    setTimelineNote(
                      e.target.value,
                    )
                  }
                />

              </div>


              <button
                className="primary-action"
                onClick={addNote}
              >
                Add Note
              </button>

            </div>


            {/* =================================================
                TIMELINE NOTES
                ================================================= */}

            <div className="timeline-list">

              {caseNotes.map(
                (
                  note,
                  index,
                ) => (

                  <article
                    key={index}
                    className="timeline-item"
                  >

                    <div className="timeline-marker">
                      <span />
                    </div>


                    <div className="timeline-content">

                      <div className="timeline-top">

                        <h4>
                          {note.analyst}
                        </h4>

                        <small>
                          {
                            new Date(
                              note.created_at,
                            ).toLocaleString()
                          }
                        </small>

                      </div>


                      <p>
                        {note.note}
                      </p>

                    </div>

                  </article>

                ),
              )}

            </div>

          </div>

        </section>

      </div>

    </main>
  );
}

export default UpdateCase;