import { useEffect, useState } from "react";
import api from "../services/api";

function SandboxDashboard() {

  const [jobs, setJobs] =
    useState([]);

  const loadJobs =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/sandbox-jobs",
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
            <th>Status</th>
            <th>Submitted</th>

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
                  {job.status}
                </td>

                <td>
                  {job.submitted_at}
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