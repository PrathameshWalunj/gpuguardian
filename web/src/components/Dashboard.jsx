import React, { useState, useEffect, useRef } from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

// maximum no of data points to display in graph, 1 minute of data with updates every second
const MAX_DATA_POINTS = 60; 

const Dashboard = () => {
    // state to hold the latest gpu metrics recieved from the server
    const [metrics, setMetrics] = useState({
        name: '',
        memoryUsed: 0,
        memoryTotal: 0,
        utilization: 0,
        temperature: 0,
        processCount: 0
    });
    
    // state to maintain history of metrics for plotting graphs
    const [metricsHistory, setMetricsHistory] = useState([]);
