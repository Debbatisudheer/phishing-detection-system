import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import api from "../services/api";

import LoadingSpinner
from "../components/common/LoadingSpinner";

import BackendOfflineCard
from "../components/errors/BackendOfflineCard";

import EmptyStateCard
from "../components/errors/EmptyStateCard";

function SandboxDashboard() {

  const [jobs, setJobs] =
    useState([]);

  const [loading, setLoading] =
    useState(true);

  const [error, setError] =
    useState(false);

  const navigate =
    useNavigate();

  const loadJobs =
    async () => {

      try {

        setLoading(true);

        const response =
          await api.get(
            "/api/sandbox/reports",
          );

        setJobs(
          Array.isArray(response.data)
            ? response.data
            : [],
        );

        setError(false);

      } catch (err) {

        console.error(err);

        setJobs([]);

        setError(true);

      } finally {

        setLoading(false);

      }

    };

  useEffect(() => {

    loadJobs();

    const interval =
      setInterval(
        loadJobs,
        30000,
      );

    return () =>
      clearInterval(
        interval,
      );

  }, []);

  if (loading) {

    return <LoadingSpinner />;

  }

  if (error) {

    return <BackendOfflineCard />;

  }

  if (!Array.isArray(jobs) || jobs.length === 0) {

    return (

      <EmptyStateCard
        title="No Sandbox Reports"
        message="No sandbox reports found."
      />

    );

  }

  return (

    <div
      style={{
        padding: "20px",
      }}
    >

      <h1>
        Sandbox Dashboard
      </h1>

      <button
        onClick={loadJobs}
        style={{
          marginBottom: "20px",
        }}
      >
        Refresh
      </button>

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

          {jobs.map((job) => (

            <tr key={job.id}>

              <td>{job.id}</td>

              <td>{job.file_name}</td>

              <td>{job.risk_score}</td>

              <td>{job.risk_level}</td>

              <td>{job.verdict}</td>

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

          ))}

        </tbody>

      </table>

    </div>

  );

}

export default SandboxDashboard;