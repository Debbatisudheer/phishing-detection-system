import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

import { useEffect, useState } from "react";
import api from "../services/api";

function IOCTrends() {

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
          "/api/ioc-trends",
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

  const totalIOCs =
    data.reduce(
      (
        total,
        item,
      ) => total + item.count,
      0,
    );

  const highestDay =
    data.length > 0
      ? data.reduce(
          (
            max,
            item,
          ) =>
            item.count > max.count
              ? item
              : max,
        )
      : null;

  return (

    <div
      style={{
        padding: "20px",
      }}
    >

      <h1>
        IOC Trends Dashboard
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
            Total IOC Events
          </h3>

          <h2>
            {totalIOCs}
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
            Peak Day
          </h3>

          <h4>
            {
              highestDay
                ?.date
            }
          </h4>
        </div>

      </div>

      <ResponsiveContainer
        width="100%"
        height={500}
      >

        <LineChart
          data={data}
        >

          <CartesianGrid
            strokeDasharray="3 3"
          />

          <XAxis
            dataKey="date"
          />

          <YAxis />

          <Tooltip />

          <Line
            type="monotone"
            dataKey="count"
            stroke="#00C49F"
            strokeWidth={3}
          />

        </LineChart>

      </ResponsiveContainer>

    </div>
  );
}

export default IOCTrends;