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

    // reference to store the websocket instance to maintain a persistence connection
    const ws = useRef(null);

    // handle websocket connection
    useEffect(() => {
        ws.current = new WebSocket('ws://localhost:8080/ws');
        
        // event listener to handle incoming websocket messages
        ws.current.onmessage = (event) => {
            // parse incoming JSON data
            const data = JSON.parse(event.data);
            // update latest metrics in state
            setMetrics(data);
            
            // update metrics history
            setMetricsHistory(prev => {
                // add new fata point and remove previous points beyond limit
                const newHistory = [...prev, {
                    time: new Date(data.timestamp * 1000).toLocaleTimeString(),
                    memoryUsed: data.MemoryUsed / 1024 / 1024 / 1024, 
                    utilization: data.Utilization,
                    temperature: data.Temperature
                }].slice(-MAX_DATA_POINTS);
                
                return newHistory;
            });
        };

        return () => {
            if (ws.current) {
                ws.current.close();
            }
        };
    }, []);

    return (
        <div className="p-6 max-w-7xl mx-auto">
            <h1 className="text-3xl font-bold mb-6">GPU Guardian Dashboard</h1>
            
           
            <Card className="mb-6">
                <CardHeader>
                    <CardTitle>GPU Information</CardTitle>
                </CardHeader>
                <CardContent>
                    <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                        <div>
                            <p className="text-sm text-gray-500">Model</p>
                            <p className="text-lg font-semibold">{metrics.Name || 'N/A'}</p>
                        </div>
                        <div>
                            <p className="text-sm text-gray-500">Memory</p>
                            <p className="text-lg font-semibold">
                                {(metrics.MemoryUsed / 1024 / 1024 / 1024).toFixed(2)}GB /
                                {(metrics.MemoryTotal / 1024 / 1024 / 1024).toFixed(2)}GB
                            </p>
                        </div>
                        <div>
                            <p className="text-sm text-gray-500">Temperature</p>
                            <p className="text-lg font-semibold">{metrics.Temperature}°C</p>
                        </div>
                        <div>
                            <p className="text-sm text-gray-500">Active Processes</p>
                            <p className="text-lg font-semibold">{metrics.ProcessCount}</p>
                        </div>
                    </div>
                </CardContent>
            </Card>
            
            {/* Graphs */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                {/* Memory Usage Graph */}
                <Card>
                    <CardHeader>
                        <CardTitle>Memory Usage</CardTitle>
                    </CardHeader>
                    <CardContent className="h-64">
                        <ResponsiveContainer width="100%" height="100%">
                            <LineChart data={metricsHistory}>
                                <CartesianGrid strokeDasharray="3 3" />
                                <XAxis dataKey="time" />
                                <YAxis unit="GB" />
                                <Tooltip />
                                <Legend />
                                <Line 
                                    type="monotone" 
                                    dataKey="memoryUsed" 
                                    stroke="#3b82f6" 
                                    name="Memory Used"
                                />
                            </LineChart>
                        </ResponsiveContainer>
                    </CardContent>
                </Card>

                {/* GPU Utilization Graph */}
                <Card>
                    <CardHeader>
                        <CardTitle>GPU Utilization</CardTitle>
                    </CardHeader>
                    <CardContent className="h-64">
                        <ResponsiveContainer width="100%" height="100%">
                            <LineChart data={metricsHistory}>
                                <CartesianGrid strokeDasharray="3 3" />
                                <XAxis dataKey="time" />
                                <YAxis unit="%" />
                                <Tooltip />
                                <Legend />
                                <Line 
                                    type="monotone" 
                                    dataKey="utilization" 
                                    stroke="#10b981" 
                                    name="Utilization"
                                />
                            </LineChart>
                        </ResponsiveContainer>
                    </CardContent>
                </Card>

                {/* Temperature Graph */}
                <Card>
                    <CardHeader>
                        <CardTitle>Temperature</CardTitle>
                    </CardHeader>
                    <CardContent className="h-64">
                        <ResponsiveContainer width="100%" height="100%">
                            <LineChart data={metricsHistory}>
                                <CartesianGrid strokeDasharray="3 3" />
                                <XAxis dataKey="time" />
                                <YAxis unit="°C" />
                                <Tooltip />
                                <Legend />
                                <Line 
                                    type="monotone" 
                                    dataKey="temperature" 
                                    stroke="#ef4444" 
                                    name="Temperature"
                                />
                            </LineChart>
                        </ResponsiveContainer>
                    </CardContent>
                </Card>
            </div>
        </div>
    );
};

export default Dashboard;