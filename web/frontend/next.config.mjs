/** @type {import('next').NextConfig} */
const nextConfig = {
    output: 'standalone',
    env: {
        NEXT_PUBLIC_BACKEND_URL: process.env.NODE_ENV === 'development' ? 'http://localhost:8001' : undefined,
    },
    async rewrites() {
        return [
            {
                source: '/api/:path*',
                destination: 'http://localhost:8001/api/:path*',
            },
        ]
    },
};

export default nextConfig;
