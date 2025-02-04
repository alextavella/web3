'use client'

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/app/components/ui/card'
import { ethers } from 'ethers'
import Image from 'next/image'
import React from 'react'

type SignInProps = any
type SignInState = {
  error: boolean | string
}

const SignIn: React.FunctionComponent<SignInProps> = () => {
  const [signInState, setSignInState] = React.useState<SignInState>({
    error: false,
  })

  const handleSignIn = async (wallet: string) => {
    if (!window.ethereum) {
      setSignInState({ error: `No ${wallet} wallet found` })
      console.log('MetaMask not installed; using read-only defaults.')
      return
    }

    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const account = await provider.send('eth_requestAccounts', []).then(res => {
      return Array.isArray(res) && res.length ? res?.[0] : null
    })

    if (!account) {
      setSignInState({ error: 'Wallet not found/allowed' })
      console.log('Wallet not found/allowed.')
      return
    }

    localStorage.setItem('wallet', account)

    const balance = await provider.getBalance(account)
    console.log('balance', ethers.utils.formatEther(balance))

    // const signer = provider.getSigner(account)
  }

  return (
    <div className="flex items-center justify-center h-screen">
      <Card className="max-w-[400px]">
        <CardHeader>
          <CardTitle>Conectar carteira</CardTitle>
          <CardDescription>
            Comece conectando-se com uma das carteiras abaixo. Certifique-se de
            guardar suas chaves secretas ou frase seed em um local seguro. Nunca
            as compartilhe com outras pessoas.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <button
            title="Rabby"
            onClick={() => handleSignIn('Rabby')}
            type="button"
          >
            <Image
              className="rounded"
              src="https://assets.pancakeswap.finance/web/wallets/rabby.png"
              alt="Rabby"
              width={50}
              height={50}
            />
          </button>
        </CardContent>
      </Card>
    </div>
  )
}

export default SignIn
