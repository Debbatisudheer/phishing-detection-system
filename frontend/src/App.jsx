import { useEffect, useState } from "react";

import DashboardCards from "./DashboardCards";
import RiskChart from "./RiskChart";
import AttackTimeline from "./AttackTimeline";
import ThreatIntelPanel from "./ThreatIntelPanel";

import {
  toast,
  ToastContainer,
} from "react-toastify";

import "react-toastify/dist/ReactToastify.css";

function App() {

  const [emails, setEmails] =
    useState([]);

  useEffect(() => {

    // Fetch existing emails
    fetch(
      "http://localhost:8081/emails"
    )
      .then((response) =>
        response.json()
      )
      .then((data) => {

        setEmails(data);
      });

    // WebSocket connection
    const socket = new WebSocket(
      "ws://localhost:8081/ws"
    );

    socket.onmessage = (event) => {

      const newEmail =
        JSON.parse(event.data);

      // High risk popup
      if (
        newEmail.risk_score >= 200
      ) {

        toast.error(
          `High Risk Alert: ${newEmail.subject}`
        );
      }

      console.log(
        "Realtime Detection:",
        newEmail
      );

      setEmails((prevEmails) => [

        newEmail,

        ...prevEmails,
      ]);
    };

    return () => {

      socket.close();
    };

  }, []);

  // Quarantine Action
  const quarantineEmail = async (
    id
  ) => {

    await fetch(
      `http://localhost:8081/quarantine/${id}`,
      {
        method: "POST",
      }
    );

    setEmails((prevEmails) =>

      prevEmails.map((email) =>

        email.id === id
          ? {
              ...email,
              decision:
                "QUARANTINE",
            }
          : email
      )
    );

    toast.error(
      "Email quarantined"
    );
  };

  // Add Analyst Note
  const addNote = async (
    id,
    note
  ) => {

    await fetch(
      `http://localhost:8081/add-note/${id}`,
      {
        method: "POST",

        headers: {
          "Content-Type":
            "application/json",
        },

        body: JSON.stringify({
          note,
        }),
      }
    );

    toast.success(
      "Analyst note added"
    );
  };

  return (

    <div
      style={{
        padding: "20px",
        fontFamily: "Arial",
        backgroundColor: "#111",
        minHeight: "100vh",
        color: "white",
      }}
    >

      <h1>
        SOC Dashboard
      </h1>

      <DashboardCards
        emails={emails}
      />

      <RiskChart
        emails={emails}
      />

      <AttackTimeline
        emails={emails}
      />

      <ThreatIntelPanel
        emails={emails}
      />

      <table
        border="1"
        cellPadding="10"
        style={{
          borderCollapse:
            "collapse",

          width: "100%",

          backgroundColor:
            "#1e1e1e",
        }}
      >

        <thead>

          <tr>

            <th>ID</th>

            <th>Sender</th>

            <th>Subject</th>

            <th>Risk Score</th>

            <th>Decision</th>
            <th>MITRE ATT&CK</th>

            <th>Action</th>

            <th>
              Analyst Notes
            </th>

          </tr>

        </thead>

        <tbody>

          {emails.map(
            (email, index) => (

              <tr
                key={index}
                style={{
                  backgroundColor:
                    email.risk_score >=
                    250
                      ? "#5c1a1b"
                      : email.risk_score >=
                        150
                      ? "#5a3b00"
                      : "#1e1e1e",
                }}
              >

                <td>
                  {email.id ||
                    "LIVE"}
                </td>

                <td>
                  {email.sender}
                </td>

                <td>
                  {email.subject}
                </td>

                <td>
                  {email.risk_score}
                </td>
                <td
  style={{
    color:
      email.decision ===
      "QUARANTINE"
        ? "red"
        : "lime",

    fontWeight: "bold",
  }}
>
  {email.decision}
</td>

<td>
  {email.mitre_technique}
</td>

                <td>

                  <button
                    onClick={() =>
                      quarantineEmail(
                        email.id
                      )
                    }
                  >
                    QUARANTINE
                  </button>

                </td>

                <td>

                  <input
                    type="text"
                    placeholder="Add note"
                    defaultValue={
                      email.analyst_note
                    }

                    onBlur={(e) =>
                      addNote(
                        email.id,
                        e.target.value
                      )
                    }
                  />

                </td>

              </tr>
            )
          )}

        </tbody>

      </table>

      <ToastContainer />

    </div>
  );
}

export default App;