---
title: "Chainguard社の気になる記事メモ"
emoji: "📝"
published: true
date: 2023-02-21T21:30:00+09:00
toc: true
tags: ["sbom", "chainguard"]
---

# はじめに

昨今、ソフトウェアサプライチェーンの問題に対処するため、SBOMを用いたソフトウェアサプライチェーンの管理に関する議論が活発。SBOMに関しては、2021年のサイバーセキュリティ関する米国大統領令で触れられたことをはじめ、日本の経済産業省でもSBOM活用に関する議論が続けられている。

[サイバー・フィジカル・セキュリティ確保に向けたソフトウェア管理手法等検討タスクフォース （METI/経済産業省）](https://www.meti.go.jp/shingikai/mono_info_service/sangyo_cyber/wg_seido/wg_bunyaodan/software/index.html)

ここ最近、Chainguard社がSBOM関連の興味深い記事をいくつか出しているので、備忘録がてらメモする。

## (2022/11/09) [Software Dark Matter is the Enemy of Software Transparency](https://www.chainguard.dev/unchained/software-dark-matter-is-the-enemy-of-software-transparency)

宇宙の27%は観測不可能なdark matterで構成されているが、数百のポピュラーなOSSのコンテナを調べたところ、ソフトウェアにもパッケージマネージャーやSBOMなどで捕捉されないパッケージなどの **ソフトウェアダークマター** が32%ほど含まれていたという話。**ソフトウェアダークマター** は、記事内ではOSのパッケージマネージャー (`apt` や `apk`)で管理されていないファイルとして定義されている。

**ソフトウェアダークマター** の現実世界での問題点として、ソフトウェア解析ツールがソフトウェアコンポーネントを正しく特定できないためにセキュリティスキャナーが検出できない脆弱性があることや、SBOMによる情報管理が不完全になる点などを挙げている。

筆者は、<u>350のポピュラーなOSSのコンテナ内のファイルの内、どの程度の割合が **ソフトウェアダークマター** なのかを自作ツール ([chainguard-dev/darkfiles](https://github.com/chainguard-dev/darkfiles)) で調査</u>している。結果は、<u>平均32%のファイルが **ソフトウェアダークマター** </u>であり、調査対象のコンテナ内に含まれるファイル数で重みづけを行うと、その割合は 63% まで上昇した。また、調査したコンテナの約30%は **ソフトウェアダークマター** が1%未満であったため、ソフトウェアダークマターを少なくすることもできること示している。

(実際に計測を行っているのは以下のファイルの処理だと思うが、少し大味な測定方法な気もする (ちゃんと読めてないだけかもしれないが))

[https://github.com/chainguard-dev/darkfiles/blob/main/pkg/distro/distro.go](https://github.com/chainguard-dev/darkfiles/blob/main/pkg/distro/distro.go)

---

調査結果のソフトウェアダークマターのパーセンテージは少し誇張されている気もするが、ソフトウェアダークマターという言葉はキャッチーであるし、現実に問題となる可能性の高い事象であるため、言語化されたことで改善が進んでいくと良いなと思った。


## (2022/12/21) [Are SBOMs Any Good? Preliminary Measurement of the Quality of Open Source Project SBOMs](https://www.chainguard.dev/unchained/are-sboms-any-good-preliminary-measurement-of-the-quality-of-open-source-project-sboms)

SBOMの質が悪いとSBOMの情報を利用するツールの結果も悪化するという問題意識から、OSSプロジェクトの公開されているSBOMの質的調査を実施したという記事。収集したSBOMは、GitHub ([chainguard-dev/bom-shelter](https://github.com/chainguard-dev/bom-shelter)) で公開している。

収集したSBOMは、[Sourcegraph](https://sourcegraph.com/search) を使ってある特定の名前 (例: `spdx.json`)のファイルを検索して見つけたもの。SBOMの質的評価は [eBay/sbom-scorecard](https://github.com/eBay/sbom-scorecard) と [spdx/ntia-conformance-checker](https://github.com/spdx/ntia-conformance-checker) を使用して行っている。

結果としては、調査したSBOMの内、<u>80% はライセンスに関する情報が欠落</u>しており、<u>40% はパッケージのバージョン情報が欠落</u>していた。また、SPDXコミュニティのNTIA最小要素チェックツール ([spdx/ntia-conformance-checker](https://github.com/spdx/ntia-conformance-checker)) の結果より、<u>今回収集したSBOMには、NTIA最小要素に適合するSBOMが無い</u>ことがわかった。(NTIA最小要素に関する言及としては、次に紹介する記事でより詳細に調査をしている)

---

SBOMに書かれている情報が、本当に対象のソフトウェアのすべての構成物を表しているかは保証がなく、この調査では実際に不完全なSBOMファイルがSBOMとして公開されていることを示している。ただ、個人的な考えだが、SBOMファイルがリポジトリにコミットして管理されるというのはあまり正しい管理方法だとは思えず、その視点で見ると今回の収集方法で集めたSBOMファイルはそもそも管理方法自体に問題がある気がするので、低い品質のSBOMを集めたことがこの結果につながったのでは無いかという疑念が残る。(ただ、次の記事でより範囲の広い詳細な調査を行なっているため、この疑念は今では晴れている)

## (2023/01/19) [Are SBOMs Good Enough for Government Work?](https://www.chainguard.dev/unchained/are-sboms-good-enough-for-government-work)

一つ前で紹介した記事と同様、SBOMの質を評価する記事。

DockerHubで公開されているPopularなコンテナイメージからSBOMを生成し、生成された約3,000のSBOMファイルが、[NTIAの最小要素]([https://ntia.gov/sites/default/files/publications/sbom_minimum_elements_report_0.pdf](https://ntia.gov/sites/default/files/publications/sbom_minimum_elements_report_0.pdf)) を満たしているかを検証している。NTIA最小要素への適合は、SPDXコミュニティにより開発が行われている [spdx/ntia-conformance-checker](https://github.com/spdx/ntia-conformance-checker) を使用。

対象となるSBOMは、<u>Docker Hubで公開されている1,000個のコンテナイメージから、様々なSBOM生成ツール ([syft](https://github.com/anchore/syft), [bom](https://github.com/kubernetes-sigs/bom), [Trivy](https://github.com/aquasecurity/trivy), [Tern](https://github.com/tern-tools/tern)) を使用して生成</u>したもの。(調査用のコードは [GitHub](https://github.com/chainguard-dev/bom-shelter/tree/main/in-the-lab/spdx-popular-containers) で公開されている)

結果としては、調査したSBOMの内、<u>NTIA最小要素を完全に満たすSBOMはたったの1%</u>だった。ただ、この結果は、Supplierに関する項目がすべてのコンポーネントにおいて満たされていたSBOMが1%しかなかったということに影響を大きく受けているため、Supplierに関する項目を除外すれば数値はもっと上がるものと思われる。SBOM生成ツールごとに比較では、syftのみがNTIA最小要素準拠のSBOMを生成できていた。また、bomも`Author`、`Supplier`以外の項目では高い適合率のSBOMを生成できていた。

---

Supplierという項目は、対象がOSSであると考えるとパッケージマネージャーの情報が入ることが多いと思うが、ビルド後の成果物からは取得先の情報は失われていることが多いと考えられるため、確かにこの項目の扱いは難しいように感じた。

## (2023/01/26) [Make SBOMs, not GuessBOMs: Why we need to shift left on SBOM generation](https://www.chainguard.dev/unchained/make-sboms-not-guessboms-why-we-need-to-shift-left-on-sbom-generation)

SBOMは、ビルド後の後処理としてSCAツールを実行することで生成する手法が主流だが、その方法では問題がある。  
一つ目に挙げられている問題は、例えば、Golangのコンテナイメージから生成されたSBOMにはGolang自体が含まれていないというように、SCAツールでは自分自身のことを認識できない場合がある。(記事から参照されている以下のTweetではPythonイメージの話をしているが)

{{< tweet user="lorenc_dan" id="1563896969021558784" >}}

その他にも、解析対象の中にはSCAツールでは発見できない資産 (**Software dark matter**) が30~60%の割合で存在しているという問題もある。これらの問題を含むSBOMは、疑わしい情報である `GuessBOM` なのでは? という記事。

これを解決する方法として、NTIAも推奨しているように、ビルド後の後処理としてSCAツールを実行するのではなく、ビルド時の動作をキャプチャしながらSBOMに含める情報を集めるのが良いのではないかと提言している。ビルド時であれば、ビルド後の資産からは消えてしまう可能性のあるソースコードなどの元の情報にアクセスすることもできる。

しかし、現時点ではビルド時の情報を集めてSBOM を生成するツールはあまりなく、それが実現できるツールとしてChainguard社が開発している [`apko`](https://github.com/chainguard-dev/apko) を紹介している。 `apko` ではコンテナイメージのビルドを実行することで同時にSBOMファイルを生成することができ、さらに生成されたSBOMに対するSigstoreによる署名を行うこともできるらしい。

GuessBOMである可能性の高い成果物スキャンにより生成されたSBOMではなく、生成物に含まれる署名付きのSBOMを利用することで、SBOM内の情報の信頼性が高まるという話。

---

`apko`の宣伝エントリという感じもするが、語られている問題意識としては非常に納得できるものだった。
