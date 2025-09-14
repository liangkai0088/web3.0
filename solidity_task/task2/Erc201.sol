//SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/**
homework:
参考 openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的 ERC20 代币合约。要求：
合约包含以下标准 ERC20 功能：
    balanceOf：查询账户余额。
    transfer：转账。
    approve 和 transferFrom：授权和代扣转账。
    使用 event 记录转账和授权操作。
    提供 mint 函数，允许合约所有者增发代币。
提示：
    使用 mapping 存储账户余额和授权信息。
    使用 event 定义 Transfer 和 Approval 事件。
    部署到sepolia 测试网，导入到自己的钱包
 */

 contract Erc201{
    event Transfer(address indexed from,address indexed to,uint256 value);
    event Approval(address indexed owner,address indexed spender,uint256 value);

    //名称
    string private _name;
    //符号
    string private _symbol;
    //精度
    uint8 private _decimals;
    //总量
    uint256 private _totalSupply;
    //余额
    mapping(address => uint256) private _balances;
    //合约拥有者
    address private _owner;
    //授权
    mapping(address => mapping(address => uint256)) private _allowances;

    constructor(string memory name_,string memory symbol_){ 
        _name = name_;      
        _symbol = symbol_;
        _decimals = 18;
        _totalSupply = 1000000;
        _owner = msg.sender;
        //创建代币
        mint(_owner,_totalSupply * 10 ** _decimals);
    }
    function name() public view returns(string memory){
        return _name;
    }
    function symbol() public view returns(string memory){
        return _symbol;
    }
    function decimals() public view returns(uint8){
        return _decimals;
    }
    function totalSupply() public view returns(uint256){
        return _totalSupply;
    }
    //查询余额
    function balanceOf(address account) public view returns(uint256){
        return _balances[account];
    }


    //授权
    function transfer(address to,uint256 amount) public returns(bool){ 
        return _transfer(msg.sender,to,amount);
    }
    function _transfer(address from,address to,uint256 amount) internal returns(bool){
        require(from != address(0),"ERC20: transfer from the zero address");
        require(to != address(0),"ERC20: transfer to the zero address");
        require(amount > 0,"ERC20: transfer amount must be greater than 0");
        require(_balances[from] >= amount,"ERC20: transfer amount exceeds balance");
        _balances[from] -= amount;
        _balances[to] += amount;
        emit Transfer(from,to,amount);
        return true;
    }

    //授权
    function approve(address owner ,address spender,uint256 amount) public returns(bool){
        require(spender != address(0),"ERC20: approve spender the zero address");
        require(owner != address(0),"ERC20: approve owner the zero address");
        require(_balances[owner] >= amount,"ERC20: approve amount must be greater than balance");
        _allowances[owner][spender] = amount;
        emit Approval(owner,spender,amount);
        return true;
    }

    //代扣
    function transferFrom(address from,address to,uint256 amount) public returns(bool){
        bool res = approve(from, to, _allowances[from][msg.sender] - amount);
        if(res){
            return _transfer(from,to,amount);
        }
        return false;
    }

    function mint(address to,uint256 amount) public returns(bool){
        require(msg.sender == _owner, "ERC20Token: only owner allowed");
        require(to != address(0),"ERC20: mint to the zero address");
        //增加发行量xing 的代币
        _totalSupply += amount;
        //增加账户余额
        _balances[to] += amount;
        //触发事件
        emit Transfer(address(0),to,amount);
        return true;
    }
 }