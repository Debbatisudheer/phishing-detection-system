import { useState } from "react";
import api from "../services/api";

function InvestigationWorkbench() {

  const [reputation, setReputation] =
  useState(null);
  
  const [ioc, setIOC] =
    useState("");

  const [result, setResult] =
    useState(null);

  const [notes, setNotes] =
    useState([]);

  const [analyst, setAnalyst] =
    useState("");

  const [noteText, setNoteText] =
    useState("");

  const loadNotes =
    async (iocValue) => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          `/api/get-notes?ioc=${encodeURIComponent(iocValue)}`,
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setNotes(
        response.data,
      );
    };

    const loadReputation =
  async (iocValue) => {

    const token =
      localStorage.getItem(
        "token",
      );

    const response =
      await api.get(
        `/api/ioc-reputation?ioc=${encodeURIComponent(iocValue)}`,
        {
          headers: {
            Authorization:
              `Bearer ${token}`,
          },
        },
      );

    setReputation(
      response.data,
    );
  };

  const investigate =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          `/api/investigation-summary?ioc=${encodeURIComponent(ioc)}`,
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setResult(
        response.data,
      );

      loadNotes(
        ioc,
      );

      loadReputation(
  ioc,
);
    };

  const saveNote =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      await api.post(
        "/api/notes",
        {
          ioc:
            result.ioc,
          analyst,
          notes:
            noteText,
        },
        {
          headers: {
            Authorization:
              `Bearer ${token}`,
          },
        },
      );

      loadNotes(
        result.ioc,
      );

      setNoteText("");
    };

  return (

    <div
      style={{
        padding: "20px",
      }}
    >

      <h1>
        Investigation Workbench
      </h1>

      <input
        value={ioc}
        onChange={(e) =>
          setIOC(
            e.target.value,
          )
        }
        placeholder="Enter IOC"
        style={{
          width: "500px",
        }}
      />

      <button
        onClick={investigate}
      >
        Investigate
      </button>

      {result && (

        <div
          style={{
            marginTop: "30px",
          }}
        >

          <h2>
            IOC Details
          </h2>

          <p>
            <b>IOC:</b>{" "}
            {result.ioc}
          </p>

          <p>
            <b>Occurrences:</b>{" "}
            {result.count}
          </p>

          <p>
            <b>Risk Level:</b>{" "}
            {result.risk_level}
          </p>

          <p>
            <b>Verdict:</b>{" "}
            {result.verdict}
          </p>

          <p>
            <b>MITRE:</b>{" "}
            {result.mitre}
          </p>

          <p>
  <b>Reputation:</b>
  {" "}
  {reputation?.reputation}
</p>

<p>
  <b>Threat Intel Source:</b>
  {" "}
  {reputation?.source}
</p>

          <p>
            <b>First Seen:</b>{" "}
            {result.first_seen}
          </p>

          <p>
            <b>Last Seen:</b>{" "}
            {result.last_seen}
          </p>

          <h3>
            Sources
          </h3>

          <ul>

            {result.sources?.map(
              (
                source,
                index,
              ) => (

                <li key={index}>
                  {source}
                </li>
              ),
            )}

          </ul>

          <h3>
            Files
          </h3>

          <ul>

            {result.files?.map(
              (
                file,
                index,
              ) => (

                <li key={index}>
                  {file}
                </li>
              ),
            )}

          </ul>

          <hr />

          <h3>
            Analyst Notes
          </h3>

          <input
            placeholder="Analyst Name"
            value={analyst}
            onChange={(e) =>
              setAnalyst(
                e.target.value,
              )
            }
          />

          <br />
          <br />

          <textarea
            rows="4"
            cols="60"
            placeholder="Enter note"
            value={noteText}
            onChange={(e) =>
              setNoteText(
                e.target.value,
              )
            }
          />

          <br />
          <br />

          <button
            onClick={saveNote}
          >
            Save Note
          </button>

          <hr />

          <h3>
            Previous Notes
          </h3>

          {notes?.map(
            (
              note,
              index,
            ) => (

              <div
                key={index}
                style={{
                  border:
                    "1px solid gray",
                  padding:
                    "10px",
                  marginBottom:
                    "10px",
                }}
              >

                <b>
                  {note.analyst}
                </b>

                <p>
                  {note.notes}
                </p>

                <small>
                  {note.created_at}
                </small>

              </div>
            ),
          )}

        </div>
      )}

    </div>
  );
}

export default InvestigationWorkbench;