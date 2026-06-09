import { useState } from "react";
import api from "../services/api";

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

    <div
      style={{
        padding: "20px",
      }}
    >

      <h1>
        Update Case
      </h1>

      <input
        placeholder="Case ID"
        value={caseId}
        onChange={(e) =>
          setCaseId(
            e.target.value,
          )
        }
      />

      <button
        onClick={loadNotes}
      >
        Load Notes
      </button>

      <br />
      <br />

      <input
        placeholder="Status"
        value={status}
        onChange={(e) =>
          setStatus(
            e.target.value,
          )
        }
      />

      <br />
      <br />

      <textarea
        placeholder="Case Notes"
        value={notes}
        onChange={(e) =>
          setNotes(
            e.target.value,
          )
        }
      />

      <br />
      <br />

      <button
        onClick={updateCase}
      >
        Update Case
      </button>

      <hr />

      <h2>
        Analyst Timeline
      </h2>

      <textarea
        placeholder="Add Investigation Note"
        value={timelineNote}
        onChange={(e) =>
          setTimelineNote(
            e.target.value,
          )
        }
      />

      <br />
      <br />

      <button
        onClick={addNote}
      >
        Add Note
      </button>

      <hr />

      {caseNotes.map(
        (
          note,
          index,
        ) => (

          <div
            key={index}
            className="card"
          >

            <h4>
              {note.analyst}
            </h4>

            <p>
              {note.note}
            </p>

            <small>
  {
    new Date(
      note.created_at,
    ).toLocaleString()
  }
</small>

          </div>
        ),
      )}

    </div>
  );
}

export default UpdateCase;