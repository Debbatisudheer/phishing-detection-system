import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../services/api";

function SandboxDashboard() {

  const [jobs, setJobs] =
    useState([]);

  const navigate =
    useNavigate();

  const loadJobs =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/sandbox/reports",
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setJobs(
        response.data,
      );
    };

  useEffect(() => {

    loadJobs();

  }, []);

  return (

    <div
      style={{
        padding: "20px",
      }}
    >

      <h1>
        Sandbox Dashboard
      </h1>

      <table
        border="1"
        cellPadding="10"
      >

        <thead>

          <tr>

            <th>ID</th>
            <th>File</th>
            <th>Risk Score</th>
            <th>Risk Level</th>
            <th>Verdict</th>
            <th>Action</th>

          </tr>

        </thead>

        <tbody>

          {jobs.map(
            (job) => (

              <tr key={job.id}>

                <td>
                  {job.id}
                </td>

                <td>
                  {job.file_name}
                </td>

                <td>
                  {job.risk_score}
                </td>

                <td>
                  {job.risk_level}
                </td>

                <td>
                  {job.verdict}
                </td>

                <td>

                  <button
                    onClick={() =>
                      navigate(
                        `/sandbox-report/${job.id}`,
                      )
                    }
                  >
                    View
                  </button>

                </td>

              </tr>
            ),
          )}

        </tbody>

      </table>

    </div>
  );
}

export default SandboxDashboard;