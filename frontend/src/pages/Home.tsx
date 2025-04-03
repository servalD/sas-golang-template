import { Rocket } from 'lucide-react';

export default function Home() {
  return (
    <div className="min-h-screen bg-gradient-to-b from-gray-900 to-gray-800 text-white">
      <div className="container mx-auto px-4 py-16">
        <div className="flex flex-col items-center text-center">
          <Rocket className="w-16 h-16 mb-8 text-blue-400" />
          <h1 className="text-5xl font-bold mb-6">Bienvenue sur NFT Explorer</h1>
          <p className="text-xl text-gray-300 max-w-2xl mb-8">
            Découvrez et explorez les collections NFT les plus exceptionnelles. 
            Une plateforme unique pour les passionnés d'art numérique.
          </p>
          <img
            src="https://images.unsplash.com/photo-1620641788421-7a1c342ea42e?ixlib=rb-1.2.1&auto=format&fit=crop&w=1200&q=80"
            alt="NFT Collection"
            className="rounded-lg shadow-2xl max-w-4xl w-full"
          />
        </div>
      </div>
    </div>
  );
}