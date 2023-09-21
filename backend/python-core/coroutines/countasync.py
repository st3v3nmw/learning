import asyncio
import time


async def count():
    print("One")
    await asyncio.sleep(1)
    print("Two")

async def main():
    await asyncio.gather(count(), count(), count())

start = time.perf_counter()
asyncio.run(main())
elapsed = time.perf_counter() - start
print(f"Executed in {elapsed:.2f} seconds")
