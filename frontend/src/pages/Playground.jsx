import { useState, useEffect } from "react";

import SampleList from "../components/playground/SampleList";
import EmailEditor from "../components/playground/EmailEditor";
import ResultsPanel from "../components/playground/ResultsPanel";

import "../components/playground/Playground.css";

import api from "../services/api";
import loadDemoFile from "../components/playground/DemoFileLoader";

function Playground() {

  const [selectedEmail, setSelectedEmail] =
    useState(null);

  const [results, setResults] =
    useState(null);

  const [loading, setLoading] =
    useState(false);

  const [sandboxJobID, setSandboxJobID] =
    useState(null);

const [sandboxStatus, setSandboxStatus] =
    useState(null);

  const handleSelectSample =
    (sample) => {

      setSelectedEmail(
        sample,
      );

      setResults(
        null,
      );

    };

  const analyzeEmail =
    async (email) => {

      try {

        setLoading(
          true,
        );

        

        // -------------------------
        // DEMO FILE ANALYSIS
        // -------------------------

        if (
          email.demoFile
        ) {

         const result =
    await loadDemoFile(
        email.demoFile,
    );

    console.log(result);

setResults(result);

setSandboxJobID(
    result.sandbox_job_id,
);

setSandboxStatus(
    result.sandbox_status,
);

console.log(
    "JOB:",
    result.sandbox_job_id,
);

console.log(
    "STATUS:",
    result.sandbox_status,
);

return;

setResults(
    result,
);

setSandboxJobID(
    result.sandbox_job_id,
);

setSandboxStatus(
    result.sandbox_status,
);

return;

        }

        // -------------------------
        // EMAIL ANALYSIS
        // -------------------------

        const response =
          await api.post(
            "/api/analyze-email",
            {
              subject:
                email.subject,
              body:
                email.body,
            },
          );

        setResults(
          response.data,
        );



        setSandboxJobID(
    response.data.sandbox_job_id,
);

setSandboxStatus(
    response.data.sandbox_status,
);


      } catch (err) {

        console.error(
          err,
        );

      } finally {

        setLoading(
          false,
        );

      }

    };

    useEffect(() => {

    if (
        !sandboxJobID ||
        sandboxStatus !== "RUNNING"
    ) {

        return;

    }

    const interval = setInterval(
        async () => {

            try {

                console.log(
                    "Polling Sandbox:",
                    sandboxJobID,
                );

                const response =
                    await api.get(
                        `/api/sandbox-report/${sandboxJobID}`,
                    );

                console.log(
                    "Sandbox Completed:",
                    response.data,
                );

                setResults(
                    previous => ({

                        ...previous,

                        sandbox_completed: true,

                        sandbox_status: "COMPLETED",

                        sandbox: response.data,

                    }),
                );

                setSandboxStatus(
                    "COMPLETED",
                );

                clearInterval(
                    interval,
                );

            } catch {

                // still running

            }

        },

        2000,

    );

    return () =>
        clearInterval(
            interval,
        );

}, [

    sandboxJobID,

    sandboxStatus,

]);

  return (

    <div
      className="playground-container"
    >

      <SampleList
        onSelect={
          handleSelectSample
        }
      />

      <EmailEditor
        selectedEmail={
          selectedEmail
        }
        onAnalyze={
          analyzeEmail
        }
      />

      <ResultsPanel
        results={
          results
        }
        loading={
          loading
        }
      />

    </div>

  );

}

export default Playground;