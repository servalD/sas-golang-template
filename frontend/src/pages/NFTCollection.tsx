import { useState, useEffect } from 'react';
import axios from 'axios';

interface NFT {
  id: string;
  name: string;
  image: string;
  description: string;
}

export default function NFTCollection() {
  const [nfts, setNfts] = useState<NFT[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchNFTs = async () => {
      try {
        const response = await axios.get('https://api.example.com/nfts');
        setNfts(response.data);
        setLoading(false);
      } catch (err) {
        setError('Erreur lors du chargement de la collection NFT');
        setLoading(false);
      }
    };

    fetchNFTs();
  }, []);

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded">
          {error}
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100 py-12 px-4">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-4xl font-bold text-center mb-12">Collection NFT</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {nfts.map((nft) => (
            <div key={nft.id} className="bg-white rounded-lg shadow-md overflow-hidden">
              <img src={nft.image} alt={nft.name} className="w-full h-64 object-cover" />
              <div className="p-6">
                <h2 className="text-xl font-bold mb-2">{nft.name}</h2>
                <p className="text-gray-600">{nft.description}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}