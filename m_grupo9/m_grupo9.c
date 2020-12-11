#include <linux/init.h>
#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/fs.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <asm/uaccess.h>
#include <linux/hugetlb.h>
#include <linux/mm.h>
#include <linux/mman.h>
#include <linux/mmzone.h>
#include <linux/syscalls.h>
#include <linux/swap.h>
#include <linux/swapfile.h>
#include <linux/vmstat.h>
#include <linux/atomic.h>


struct sysinfo info;

static int leer_memoria(struct seq_file *m, void *v){

    #define Convert(x) ((x) << (PAGE_SHIFT - 10))
	si_meminfo(&info); 
	seq_printf(m, "{\n");
	seq_printf(m, "\"Mem_Total\":%8lu,\n",Convert(info.totalram)/1024);
	seq_printf(m, "\"Mem_Libre\":%8lu,\n",(Convert(info.freeram))/1024);
	seq_printf(m, "\"Buffer\":%8lu,\n",(Convert(info.bufferram))/1024);
	seq_printf(m, "\"Compartida\":%8lu\n",(Convert(info.sharedram))/1024);
	seq_printf(m, "}\n");
	#undef K
	return 0;

}

static int mem_info_open(struct inode *inode, struct file *file){
	return single_open(file, leer_memoria, NULL);
}

static const struct file_operations mem_info_fops = {
	.owner = THIS_MODULE,
	.open = mem_info_open,
	.read = seq_read,
	.llseek = seq_lseek,
	.release = single_release,
};

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de Memoria - Sistemas Operativos 2");

static int __init m_grupo9_init(void)
{
	printk(KERN_INFO "Buenas, att: Grupo 9, monitor de memoria\n");
	proc_create("m_grupo9", 0, NULL, &mem_info_fops);
	return 0;
}

static void __exit m_grupo9_cleanup(void)
{
	remove_proc_entry("m_grupo9", NULL);
	printk(KERN_INFO "Bai, att: Grupo 9 y este fue el monitor de memoria\n");
}

module_init(m_grupo9_init);
module_exit(m_grupo9_cleanup);