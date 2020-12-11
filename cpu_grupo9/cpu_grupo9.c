#include <linux/fs.h>
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/list.h>
#include <linux/module.h>
#include <linux/proc_fs.h>
#include <linux/sched.h>
#include <linux/seq_file.h>
#include <linux/slab.h>
#include <linux/string.h>
#include <linux/types.h>
#include <linux/mm.h>
#include <linux/sched/signal.h>
#include <linux/cred.h>



char buffer[256];
struct sysinfo info;
struct task_struct *task_list;
struct task_struct *task_child;
struct list_head *list;

char * get_task_state(long state)
{
    switch (state) {
        case TASK_RUNNING:
            return "Ejecucion";
        case TASK_INTERRUPTIBLE:
            return "Suspendido";
        case TASK_UNINTERRUPTIBLE:
            return "Suspendido";
        case __TASK_STOPPED:
            return "Detenido";
        case __TASK_TRACED:
            return "Detenido";
        case TASK_STOPPED:
            return "Detenido";
        case EXIT_ZOMBIE:
            return "Zombie";
        default:
        {
            sprintf(buffer, "Desconocido%ld", state);
            return buffer;
        }
    }
}

int get_id_task_state(long state)
{
    switch (state) {
        case TASK_RUNNING:
            return 0;
        case TASK_INTERRUPTIBLE:
            return 1;
        case TASK_UNINTERRUPTIBLE:
            return 1;
        case __TASK_STOPPED:
            return 2;
        case __TASK_TRACED:
            return 2;
        case TASK_STOPPED:
            return 2;
        case EXIT_ZOMBIE:
            return 3;
        default:
        {
            return 4;
        }
    }
}

void write_process(struct seq_file *m, struct task_struct *s,long totalM){
    #define Convert(x) ((x) << (PAGE_SHIFT - 10))
    seq_printf(m,"PID: %d\t\tNombre: %s\t\tMemoria: %ld\t\tTotalM: %d\t\tEstado: %s\n",s->pid, s->comm,get_mm_rss(s->mm), pid_vnr(task_pgrp(s)), get_task_state(s->state));
    #undef K
}

void write_process_json(struct seq_file *m, struct task_struct *s){
    #define Convert(x) ((x) << (PAGE_SHIFT - 10))
    seq_printf(m,"\"PID\":\"%d\",\"Nombre\":\"%s\",\"Memoria\":\"%ld\",\"Usuario\":\"%d\",\"Estado\":\"%s\"\n",s->pid, s->comm,Convert(get_mm_rss(s->mm)),__kuid_val(s->real_cred->uid), get_task_state(s->state));
    #undef K
}

static int pstreeG(struct seq_file *m, void *v)
{
    
    unsigned int process_count = 0;
    int i = 0;
    int j = 0;
    int vec[] ={0,0,0,0,0};
    int stateTask ;
    int n = 0;
    pr_info("%s: In init\n", __func__);
    si_meminfo(&info);
    seq_printf(m,"{\n");
    seq_printf(m,"\"Lista\":[\n");
    for_each_process(task_list) {
        if(task_list->mm){
            if(i==0){
                i +=1;
            }else{
                seq_printf(m,",");
            }
            seq_printf(m,"{\n");
            write_process_json(m,task_list);
            stateTask = get_id_task_state(task_list->state);
            vec[stateTask] = vec[stateTask] + 1;
            process_count++;  
            j = 0;
            list_for_each(list, &task_list->children){   
                if(j == 0){
                    seq_printf(m,",\"hijos\":[\n");
                    j+=1;
                }else{
                    seq_printf(m,",");
                }
                seq_printf(m,"{\n");                
                task_child = list_entry( list, struct task_struct, sibling );   
                write_process_json(m,task_child);
                seq_printf(m,"}\n");
                //process_count++;   
            }
            if(j > 0){
                seq_printf(m,"]\n");
            }
            seq_printf(m,"}\n");   
        }else{
            n+=1;
        }
    }
    
    seq_printf(m,"],\n");
    #define Convert(x) ((x) << (PAGE_SHIFT - 10))
    seq_printf(m,"\"MemoriaTotal\":%ld,\n",Convert(info.totalram));
    #undef K
    seq_printf(m,"\"ProcesosTotal\":%d,\n",process_count);
    seq_printf(m,"\"ProcesosEjecucion\":%d,\n",vec[0]);
    seq_printf(m,"\"ProcesosSustendido\":%d,\n",vec[1]);
    seq_printf(m,"\"ProcesosDetenidos\":%d,\n",vec[2]);
    seq_printf(m,"\"ProcesosZombie\":%d,\n",vec[3]);
    seq_printf(m,"\"ProcesosDesconocidos\":%d\n",vec[4]);
    seq_printf(m,"}\n");
    pr_info("Number of processes:%u\n", process_count);
    pr_info("Number of processes desc:%u\n", n);
    return 0;
}

static int cpu_info_proc_open(struct inode *inode, struct file *file)
{
    return single_open(file, pstreeG, NULL);
}

static const struct file_operations cpu_info_proc_fops = {
    .open       = cpu_info_proc_open,
    .read       = seq_read,
    .llseek     = seq_lseek,
    .release    = single_release,
};

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de CPU - Sistemas Operativos 2");

static int __init cpu_grupo9_init(void)
{
	printk(KERN_INFO "Buenas, att: Grupo 9, este es el monitor de CPU\n");
	proc_create("cpu_grupo9", 0, NULL, &cpu_info_proc_fops);
	return 0;
}

static void __exit cpu_grupo9_cleanup(void)
{
	remove_proc_entry("cpu_grupo9", NULL);
	printk(KERN_INFO "Bai, att: Grupo 9 y este fue el monitor de CPU\n");
}

module_init(cpu_grupo9_init);
module_exit(cpu_grupo9_cleanup);