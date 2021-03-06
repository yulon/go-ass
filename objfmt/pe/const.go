package pe

const (
	ImageBaseUtil = 0x00400000

	imageAlignment = 0x00001000
	fileAlignment = 0x00000200

	import_descriptor_size = 20

	MachineI386 = 0x014c // x86 CPU
	MachineAMD64 = 0x8664 // x64 CPU

	image_nt_optional_hdr32_magic = 0x10b
	image_nt_optional_hdr64_magic = 0x20b

	image_file_relocs_stripped = 0x0001 // 文件中不存在重定位信息
	image_file_executable_image = 0x0002 // 文件是可执行的
	image_file_line_nums_stripped = 0x0004 // 不存在行信息
	image_file_local_syms_stripped = 0x0008 // 不存在符号信息
	image_file_aggresive_ws_trim = 0x0010 // 让操作系统强制整理工作区
	image_file_large_address_aware = 0x0020 // 应用程序可以处理大于2GB的地址空间
	 //	IMAGE_FILE_??? = 64 // 保留，留以后扩展
	image_file_bytes_reversed_lo = 0x0080 // 小尾方式
	image_file_32bit_machine = 0x0100 // 只在32位平台上运行
	image_file_debug_stripped = 0x0200 // 不包含调试信息。调试信息位于一个 .DBG 文件中
	image_file_removable_run_from_swap = 0x0400 // 如果映像在可移动媒体中，那么复制到交换文件并从交换文件中运行
	image_file_net_run_from_swap = 0x0800 // 如果映像在网络上，那么复制到交换文件并从交换文件中运行
	image_file_system = 0x1000 // 系统文件（如驱动程序），不能直接运行
	image_file_dll = 0x2000 // 这是一个 DLL 文件
	image_file_up_system_only = 0x4000 // 只能在单处理器机器中运行
	image_file_bytes_reversed_hi = 0x8000 // 大尾方式

	image_directory_entry_export = 0 // 指向导出表（IMAGE_EXPORT_DIRECTORY）
	image_directory_entry_import = 1 // 指向导入表（IMAGE_IMPORT_DESCRIPTOR 数组）
	image_directory_entry_resource = 2 // 指向资源（IMAGE_RESOURCE_DIRECTORY）
	image_directory_entry_exception = 3 // 指向异常处理表（IMAGE_RUNTIME_FUNCTION_ENTRY 数组）。CPU特定的并且基于表的异常处理。用于除x86之外的其它CPU上。
	image_directory_entry_security = 4 // 指向一个 WIN_CERTIFICATE 结构的列表，它定义在 WinTrust.H 中。不会被映射到内存中。因此，VirtualAddress 域是一个文件偏移，而不是一个RVA。
	image_directory_entry_basereloc = 5 // 指向基址重定位信息
	image_directory_entry_debug = 6 // 指向一个 IMAGE_DEBUG_DIRECTORY 结构数组，其中每个结构描述了映像的一些调试信息。早期的 Borland 链接器设置这个 IMAGE_DATA_DIRECTORY 结构的 Size 域为结构的数目，而不是字节大小。要得到 IMAGE_DEBUG_DIRECTORY 结构的数目，用 IMAGE_DEBUG_DIRECTORY 的大小除以这个 Size 域
	image_directory_entry_architecture = 7 // 指向特定架构数据，它是一个 IMAGE_ARCHITECTURE_HEADER 结构数组。不用于 x86 或 x64，但看来已用于 DEC/Compaq Alpha。
	image_directory_entry_globalptr = 8 // 在某些架构体系上 VirtualAddress 域是一个 RVA，被用来作为全局指针（gp）。不用于 x86，而用于 IA-64。Size 域没有被使用。参见2000年11月的 Under The Hood 专栏可得到关于 IA-64 gp 的更多信息
	image_directory_entry_tls = 9 // 指向线程局部存储初始化节
	image_directory_entry_load_config = 10 // 指向一个 IMAGE_LOAD_CONFIG_DIRECTORY 结构。IMAGE_LOAD_CONFIG_DIRECTORY 中的信息是特定于 Windows NT、Windows 2000 和 Windows XP 的(例如 GlobalFlag 值)。要把这个结构放到你的可执行文件中，你必须用名字 __load_config_used 定义一个全局结构，类型是 IMAGE_LOAD_CONFIG_DIRECTORY。对于非 x86 的其它体系，符号名是 _load_config_used (只有一个下划线)。如果你确实要包含一个 IMAGE_LOAD_CONFIG_DIRECTORY，那么在 C++ 中要得到正确的名字比较棘手。链接器看到的符号名必须是__load_config_used (两个下划线)。C++ 编译器会在全局符号前加一个下划线。另外，它还用类型信息修饰全局符号名
	image_directory_entry_bound_import = 11 // 指向一个 IMAGE_BOUND_IMPORT_DESCRIPTOR 结构数组，对应于这个映像绑定的每个 DLL。数组元素中的时间戳允许加载器快速判断绑定是否是新的。如果不是，加载器忽略绑定信息并且按正常方式解决导入 API
	image_directory_entry_iat = 12 // 指向第一个导入地址表（IAT）的开始位置。对应于每个被导入 DLL 的 IAT 都连续地排列在内存中。Size 域指出了所有 IAT 的总的大小。在写入导入函数的地址时加载器使用这个地址和 Size 域指定的大小临时地标记 IAT 为可读写
	image_directory_entry_delay_import = 13 // 指向延迟加载信息，它是一个 CImgDelayDescr 结构数组，定义在 Visual C++ 的头文件 DELAYIMP.H 中。延迟加载的 DLL 直到对它们中的 API 进行第一次调用发生时才会被装入。Windows 中并没有关于延迟加载 DLL 的知识，认识到这一点很重要。延迟加载的特征完全是由链接器和运行时库实现的
	image_directory_entry_comheader = 14 // 它指向可执行文件中 .NET 信息的最高级别信息，包括元数据。这个信息是一个 IMAGE_COR20_HEADER 结构

	image_subsystem_windows_gui = 2
	image_subsystem_windows_cui = 3

	image_numberof_directory_entries = 16

	image_scn_cnt_code = 0x00000020 // 节中包含代码
	image_scn_mem_execute = 0x20000000 // 节是可执行的
	image_scn_cnt_initialized_data = 0x00000040 // 节中包含已初始化数据
	image_scn_cnt_uninitialized_data = 0x00000080 // 节中包含未初始化数据
	image_scn_mem_discardable = 0x02000000 // 节可被丢弃。用于保存链接器使用的一些信息，包括.debug$节
	image_scn_mem_not_paged = 0x08000000 // 节不可被页交换，因此它总是存在于物理内存中。经常用于内核模式的驱动程序
	image_scn_mem_shared = 0x10000000 // 包含节的数据的物理内存页在所有用到这个可执行体的进程之间共享。因此，每个进程看到这个节中的数据值都是完全一样的。这对一个进程的所有实例之间共享全局变量很有用。要使一个节共享，可使用/section:name,S 链接器选项
	image_scn_mem_read = 0x40000000 // 节是可读的。几乎总是被设置
	image_scn_mem_write = 0x80000000 // 节是可写的
)
