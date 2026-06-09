import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

import { useEffect, useState } from "react";
import api from "../services/api";

function MITREHeatmap() {

  const [data, setData] =
    useState([]);

  useEffect(() => {

    loadData();

  }, []);

  const loadData =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/mitre-heatmap",
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setData(
        response.data,
      );
    };

  const totalDetections =
    data.reduce(
      (
        total,
        item,
      ) => total + item.count,
      0,
    );

  const topTechnique =
    data.length > 0
      ? data.reduce(
          (
            max,
            item,
          ) =>
            item.count > max.count
              ? item
              : max,
        ).technique
      : "N/A";

  return (

    <div
      style={{
        padding: "20px",
      }}
    >

      <h1>
        MITRE ATT&CK Heatmap
      </h1>

      <div
        style={{
          display: "flex",
          gap: "20px",
          marginBottom: "30px",
        }}
      >

        <div
          style={{
            border:
              "1px solid #333",
            padding: "20px",
            flex: 1,
          }}
        >
          <h3>
            Total Techniques
          </h3>

          <h2>
            {data.length}
          </h2>
        </div>

        <div
          style={{
            border:
              "1px solid #333",
            padding: "20px",
            flex: 1,
          }}
        >
          <h3>
            Total Detections
          </h3>

          <h2>
            {totalDetections}
          </h2>
        </div>

        <div
          style={{
            border:
              "1px solid #333",
            padding: "20px",
            flex: 2,
          }}
        >
          <h3>
            Top Technique
          </h3>

          <h4>
            {topTechnique}
          </h4>
        </div>

      </div>

      <ResponsiveContainer
        width="100%"
        height={500}
      >

        <BarChart
          layout="vertical"
          data={data}
          margin={{
            top: 20,
            right: 50,
            left: 150,
            bottom: 20,
          }}
        >

          <CartesianGrid
            strokeDasharray="3 3"
          />

          <XAxis
            type="number"
          />

          <YAxis
            type="category"
            dataKey="technique"
            width={250}
          />

          <Tooltip />

          <Bar
            dataKey="count"
            fill="#00C49F"
            label={{
              position:
                "right",
            }}
          />

        </BarChart>

      </ResponsiveContainer>

    </div>
  );
}

export default MITREHeatmap;