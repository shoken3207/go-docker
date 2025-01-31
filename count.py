import glob
from collections import defaultdict
import os

# カテゴリごとの行数を格納する辞書を初期化
category_counts = defaultdict(int)
category_files = defaultdict(list)

# すべての.goファイルを取得
files = glob.glob("**/*.go", recursive=True)

for file in files:
    # ファイルパスからカテゴリを取得
    # 例: api/user/handler.go -> user
    path_parts = file.split(os.sep)  # OSに応じたパス区切り文字で分割
    print(path_parts)
    category = "other"  # デフォルトカテゴリ
    
    # apiディレクトリ配下のファイルの場合
    if "internal" in path_parts:
        api_index = path_parts.index("internal")
        if len(path_parts) > api_index + 1:
            category = path_parts[api_index + 1]
    
    # ファイルの行数をカウント
    with open(file, 'r', encoding='utf-8') as f:
        lines = len(f.read().splitlines())
        category_counts[category] += lines
        category_files[category].append(file)

# 結果を表示
print("\n=== カテゴリごとの行数 ===")
for category, count in sorted(category_counts.items()):
    print(f"\n{category}:")
    print(f"合計行数: {count}")
    print("ファイル:")
    for file in category_files[category]:
        print(f"  - {file}")
