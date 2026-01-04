# 災害リスク判定マップ (Hazard Risk Checker)

![Go](https://img.shields.io/badge/Go-1.23-00ADD8?style=flat&logo=go)
![AWS](https://img.shields.io/badge/AWS-App%20Runner-232F3E?style=flat&logo=amazon-aws)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat&logo=docker)
![CI/CD](https://img.shields.io/badge/CI%2FCD-GitHub%20Actions-2088FF?style=flat&logo=github-actions)

地図上の任意の地点をクリックすることで、その場所の「災害リスク」を即座に判定・表示するWebサービスです。
Go言語によるバックエンド構築から、Docker/AWSを用いたモダンなインフラ構築、CI/CDの自動化までを一貫して実装しました。

**デモURL:** [ここにApp RunnerのURLを貼ってください]
*(※AWSの稼働状況により停止している場合があります)*

## 📷 スクリーンショット
![Top Page](https://via.placeholder.com/800x400?text=Please+Upload+ScreenShot)

## 🏗 アーキテクチャ構成

クラウドネイティブな構成を意識し、AWS App Runner を採用することで、コンテナの管理・運用の手間を最小限に抑えています。

```mermaid
graph LR
    User(User / Browser) -- HTTPS --> AppRunner(AWS App Runner)
    
    subgraph AWS Cloud
        AppRunner -- Pull Image --> ECR(Amazon ECR)
        AppRunner -- Run Container --> GoApp[Go App (Gin)]
    end

    subgraph CI/CD Pipeline
        Dev(Developer) -- Push --> GitHub
        GitHub -- Trigger --> Actions(GitHub Actions)
        Actions -- Build & Push --> ECR
    end
