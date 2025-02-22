import React, { useState } from 'react';
import axios from 'axios';
import { useRouter } from 'next/router';

const LoginPage: React.FC = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const router = useRouter();
    const [error, setError] = useState('');

    const handleSubmit = (event: React.FormEvent) => {
        event.preventDefault();

        axios.post('http://localhost:1323/login', { email, password }, { withCredentials: true })
            .then((response) => {
                router.push('/');
            })
            .catch((error) => {
                if (error.response.status === 401) {
                    return setError('メールアドレスまたはパスワードが間違っています');
                } else {
                    return setError('サーバーエラーが発生しました');
                }
            });
    };
    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <div className="w-full max-w-md p-8 space-y-6 bg-white rounded shadow-md">
                <h2 className="text-2xl font-bold text-center text-black">ログイン</h2>
                {error && <p className="text-red-500 text-center">{error}</p>}
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label className="block mb-1 text-sm font-medium text-gray-700">Email:</label>
                        <input
                            type="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            required
                            className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300 bg-white text-gray-900"
                        />
                    </div>
                    <div>
                        <label className="block mb-1 text-sm font-medium text-gray-700">Password:</label>
                        <input
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            required
                            className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300 bg-white text-gray-900"
                        />
                    </div>
                    <button type="submit" className="w-full px-4 py-2 text-black bg-blue-500 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-300">
                        ログイン
                    </button>
                </form>
            </div>
        </div>
    );
};

export default LoginPage;