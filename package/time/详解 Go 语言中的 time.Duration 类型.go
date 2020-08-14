/*
		长久以来，我一直抓狂于 Go 标准库中的 Time 包，我的抓狂来自于两个功能，一是捕获两个不同时间段之间间隔的毫秒数，
		二是将一个用毫秒表示的连续时间段与预先定义的时间段进行比较。这听起来很简单，没错，确实如此，但它的确让我抓狂了。

		在 Time 包中，定义有一个名为 Duration 的类型和一些辅助的常量：

		type Duration int64

		const (
			Nanosecond Duration = 1
			Microsecond = 1000 * Nanosecond
			Millisecond = 1000 * Microsecond
			Second = 1000 * Millisecond
			Minute = 60 * Second
			Hour = 60 * Minute
		)
	这些东西我可能已经看了有上千次了，但我的大脑依旧一片迷茫。我只是想比较两个时间段、恢复要持续的时间、
	比较持续时间的长短并且当预设的时间用完时做一些别的事情，但无论如何这个结构还是无法解决我的困扰。我写下了下面的测试代码，但它没有卵用：

*/