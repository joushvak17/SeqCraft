<div align="center">
  <h1>SeqCraft 🧬</h1>
  <p align="center">
    <strong>
      <a href="https://github.com/joushvak17/SeqCraft/releases"><img src="https://img.shields.io/github/v/release/joushvak17/SeqCraft" alt="Release Version"></a>
      <a href="https://github.com/joushvak17/SeqCraft/actions"><img src="https://img.shields.io/github/actions/workflow/status/joushvak17/SeqCraft/test-and-lint.yml" alt="Workflow Status"></a>
      <a href="https://goreportcard.com/report/github.com/joushvak17/SeqCraft"><img src="https://goreportcard.com/badge/github.com/joushvak17/SeqCraft" alt="Go Report Card"></a>
      <a href="https://github.com/joushvak17/SeqCraft/issues"><img src="https://img.shields.io/github/issues/joushvak17/SeqCraft" alt="Issues"></a>
      <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License: MIT"></a>
      <br>
      A CLI bioinformatics toolkit written in Go that provides researchers with efficient tools for FASTA file analysis, including sequence parsing, alignment, and structure prediction.
    </strong>
  </p>
</div>

## 📑 Table of Contents
- [Installation](#️-installation)
- [Roadmap](#️-roadmap)
- [License](#license)

## ⚙️ Installation
> [!NOTE]
> The following shows the various ways that SeqCraft can be installed (currently there is installation support only from the binary releases and building directly from the source).

### Binary Releases
Download pre-compiled binaries from the [releases page](https://github.com/joushvak17/SeqCraft/releases).

### From Source
If you want to build directly from the source then you can do the following:
```
git clone https://github.com/joushvak17/SeqCraft.git
cd SeqCraft
go build .\cmd\app\main.go
```
Then, you can call the `main.exe` output file to run the program.

## 🗺️ Roadmap
The following shows the features that hopefully will incorporated in the future:
- Sequence Translation: Translate DNA/RNA sequences to protein sequences
- Motif Search: Search for specific motifs in sequences
- Sequence Alignment: Pairwise sequence alignment and support for MSA
- Secondary Structure Prediction: Predict secondary structure of protein sequences
- Phylogenetic Analysis: Construct phylogenetic trees from sequence data

## 📄	License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
