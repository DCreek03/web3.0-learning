// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract ArtNFT is ERC721, Ownable {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIdCounter;
    
    // 存储每个NFT的元数据URI
    mapping(uint256 => string) private _tokenURIs;
    
    // 添加Ownable的构造函数参数
    constructor() ERC721("ArtNFT", "ART") Ownable(msg.sender) {}
    
    // 铸造NFT函数
    function mintNFT(address recipient, string memory _tokenURI) public onlyOwner {
        uint256 _tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        
        _safeMint(recipient, _tokenId);
        _setTokenURI(_tokenId, _tokenURI); // 调用内部函数设置URI
    }
    
    // 设置Token URI（依赖基础ERC721）
    function _setTokenURI(uint256 _tokenId, string memory _tokenURI) internal {
        require(_exists(_tokenId), "ERC721Metadata: URI set for nonexistent token");
        _tokenURIs[_tokenId] = _tokenURI;
    }
    
    // 获取Token URI（重写ERC721方法）
    function tokenURI(uint256 _tokenId) public view virtual override returns (string memory) {
        require(_exists(_tokenId), "ERC721Metadata: URI query for nonexistent token");
        return _tokenURIs[_tokenId];
    }
    
    // 基础ERC721提供的_exists函数
    function _exists(uint256 tokenId) internal view virtual returns (bool) {
        return _ownerOf(tokenId) != address(0);
    }
}