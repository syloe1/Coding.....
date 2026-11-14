#!/usr/bin/env python3
"""一键提交并推送：git add -A -> git commit -m "update" -> git push

用法:
    python push.py              # 提交信息默认是 "update"
    python push.py "feat: xxx"  # 自定义提交信息
"""

import subprocess
import sys
from pathlib import Path

REPO_ROOT = Path(__file__).resolve().parent
DEFAULT_MESSAGE = "update"


def run(cmd: list[str], check: bool = True) -> subprocess.CompletedProcess:
    """在仓库根目录执行命令，check=True 时非零退出码直接终止脚本。"""
    print(f"$ {' '.join(cmd)}", flush=True)
    result = subprocess.run(
        cmd,
        cwd=REPO_ROOT,
        text=True,
        encoding="utf-8",
        errors="replace",
    )
    if check and result.returncode != 0:
        print(f"\n命令执行失败（退出码 {result.returncode}），已中止。", file=sys.stderr)
        sys.exit(result.returncode)
    return result


def main() -> None:
    message = sys.argv[1] if len(sys.argv) > 1 else DEFAULT_MESSAGE

    # 确认脚本所在目录就是一个 git 仓库
    probe = subprocess.run(
        ["git", "rev-parse", "--is-inside-work-tree"],
        cwd=REPO_ROOT,
        stdout=subprocess.DEVNULL,
        stderr=subprocess.DEVNULL,
    )
    if probe.returncode != 0:
        sys.exit(f"{REPO_ROOT} 不是一个 git 仓库。")

    # 1. 暂存所有改动（新增、修改、删除）
    run(["git", "add", "-A"])

    # 2. 有改动才提交
    status = subprocess.run(
        ["git", "status", "--porcelain"],
        cwd=REPO_ROOT,
        text=True,
        encoding="utf-8",
        errors="replace",
        capture_output=True,
    )
    if status.stdout.strip():
        run(["git", "commit", "-m", message])
    else:
        print("工作区没有需要提交的改动，跳过 commit。")

    # 3. 推送到远端
    run(["git", "push"])
    print("\n完成 ✅")


if __name__ == "__main__":
    # Windows 控制台默认可能是 GBK，避免中文输出报错
    for stream in (sys.stdout, sys.stderr):
        try:
            stream.reconfigure(encoding="utf-8")
        except (AttributeError, ValueError):
            pass
    main()
